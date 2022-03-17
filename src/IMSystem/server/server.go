package main

import (
	"fmt"
	"net"
	"runtime"
	"sync"
	"time"
)

type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User
	MapLock   sync.RWMutex
	Message   chan string
}

//创建一个服务
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//监听消息广播
func (this *Server) listenMsg() {
	for {
		msg := <-this.Message
		this.MapLock.Lock()
		for _, clic := range this.OnlineMap {
			clic.C <- msg
		}
		this.MapLock.Unlock()
	}
}

//用户消息广播
func (this *Server) BroadCast(user *User, msg string) {

	sendMsg := "[" + user.Addr + "]" + user.Name + "：" + msg
	this.Message <- sendMsg
}

//获取上线信息
func (this *Server) Handler(conn net.Conn) {
	user := NewUser(conn, this)
	//用户上线
	user.OnlineHandler()
	fmt.Printf("%s connect successfully... \n", user.Name)
	//聊天
	isLive := make(chan bool)
	go user.Chat(isLive)

	this.Expire(isLive, user)
}

func (this *Server) Expire(isLive chan bool, user *User) {
	for {
		select {
		case <-isLive:
			//当前用户活跃，重置定时器
			//在执行后，会执行下一个case
		case <-time.After(time.Second * 120):

			user.privateMsg("你已超时，准备强制下线...")
			time.Sleep(500 * time.Millisecond)
			//关闭连接
			close(user.C)
			user.conn.Close()
			//清除数据
			this.MapLock.Lock()
			delete(this.OnlineMap, user.Name)
			this.MapLock.Unlock()
			//关闭线程
			runtime.Goexit()
		}
	}
}

//启动服务
func (this *Server) start() {
	//连接请求
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("Listen err...")
	}
	//关闭连接
	defer listen.Close()
	//监听消息
	go this.listenMsg()

	for {
		//获取连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err...")
			continue
		}
		go this.Handler(conn)
	}
}
