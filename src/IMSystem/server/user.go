package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
)

//json.Marshal不能导出chan类型、函数类型、 complex 类型
type User struct {
	Name   string      `json:"name"`
	Addr   string      `json:"addr"`
	C      chan string `json:"-"` //json过滤
	conn   net.Conn    `json:"-"`
	server *Server     `json:"-"`
}

//创建一个用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	//启动监听消息
	go user.ListenMsg()
	return user
}

//监听当前user channel的方法，一旦有消息 直接发给客户端
func (this *User) ListenMsg() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

//获取上线信息
func (this *User) OnlineHandler() {
	//用户上线
	this.server.MapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.MapLock.Unlock()
	this.server.BroadCast(this, "上线")
}

//获取离线信息
func (this *User) offlineHandler() {
	this.server.MapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.MapLock.Unlock()
	this.server.BroadCast(this, "离线")
	fmt.Printf("%s disconnect successfully... \n", this.Name)
}

//聊天
func (this *User) Chat(isLive chan bool) {
	buff := make([]byte, 1024)

	for {
		//阻塞、IO多路复用
		read, err := this.conn.Read(buff)
		//离线
		if read == 0 {
			this.offlineHandler()
			return
		}
		//数据读取错误
		if err != nil && err != io.EOF {
			fmt.Println("conn.Read error", err)
			return
		}
		//提取用户信息，并剔除buff末尾的/n符号
		msg := string(buff[:read-1])
		//消息处理
		this.msgHandle(msg)
		isLive <- true
	}
}

//当前用户消息
func (this *User) privateMsg(msg string) {
	this.C <- msg
}

//私聊
func (this *User) ptpMsg(name string, msg string) {
	if len(name) == 0 {
		this.privateMsg("用户不能为空")
		return
	}
	if len(msg) == 0 {
		this.privateMsg("消息内容不能为空")
		return
	}
	user, ok := this.server.OnlineMap[name]
	if !ok {
		this.privateMsg("当前用户未上线或不存在")
		return

	}
	user.privateMsg(this.Name + "私聊您：" + msg)
}

//消息处理
func (this *User) msgHandle(msg string) {
	//特殊消息处理
	if msg == "online" {
		//查询用户在线
		this.checkOnline()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//更该用户姓名
		split := strings.Split(msg, "|")
		this.rename(split[1])
	} else if len(msg) > 4 && msg[:4] == "get|" {
		//更该用户姓名
		split := strings.Split(msg, "|")
		this.getInfo(split[1])
	} else if len(msg) > 3 && msg[:3] == "to|" {
		//私聊
		split := strings.Split(msg, "|")
		if len(split) < 3 {
			this.privateMsg("消息格式错误,请使用\"to|张三|你好！\"的格式重新发送")
			return
		}
		this.ptpMsg(split[1], split[2])
	} else {
		//消息广播
		this.server.BroadCast(this, msg)
	}
}

//查询用户在线
func (this *User) checkOnline() {
	this.server.MapLock.Lock()
	for _, user := range this.server.OnlineMap {
		onlineMsg := "[" + user.Addr + "]" + user.Name + "：在线..."
		//给当前用户发送消息
		this.privateMsg(onlineMsg)
	}
	this.server.MapLock.Unlock()
}

//更改用户姓名
func (this *User) rename(newName string) {
	if len(newName) == 0 {
		this.privateMsg("名称不允许为空")
		return
	}
	_, ok := this.server.OnlineMap[newName]
	if ok {
		this.privateMsg("当前用户已存在")
	} else {
		this.server.MapLock.Lock()
		//删除当前用户信息
		delete(this.server.OnlineMap, this.Name)
		//更新用户信息
		this.Name = newName
		this.server.OnlineMap[newName] = this
		this.privateMsg("用户姓名更新成功：" + this.Name)
		this.server.MapLock.Unlock()
	}
}

//更改用户姓名
func (this *User) getInfo(newName string) {
	if len(newName) == 0 {
		this.privateMsg("名称不允许为空")
		return
	}
	user, ok := this.server.OnlineMap[newName]
	if ok {
		userJson, err := json.Marshal(user)
		if err != nil {
			fmt.Println("获取用户信息格式转换失败", err)
			this.privateMsg("获取用户信息失败")
			return
		}
		this.privateMsg("获取用户信息成功：" + string(userJson))
	}
}
