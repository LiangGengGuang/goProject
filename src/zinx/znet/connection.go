package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

/*
	Connection链接
*/
type Connection struct {

	//当前链接的socket
	Conn *net.TCPConn
	//当前链接ID
	ConnID uint32
	//当前链接是否关闭
	isClose bool
	//告知当前链接已经停止/退出
	ExitChan chan bool
	//该链接处理Router
	Router ziface.IRouter
}

// NewConnection 初始化链路模块方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) ziface.IConnection {

	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClose:  false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
}

//读取链接请求数据
func (c *Connection) startReader() {

	defer fmt.Println("connID =", c.ConnID, "reader is exit,remote is = ", c.RemoteAddr().String())
	defer c.Stop()

	for {

		//创建一个拆包对象
		dp := NewDataPackage()

		//读取id和消息长度
		headData := make([]byte, dp.GetHeadLength())
		if _, err := io.ReadFull(c.GetTCPConn(), headData); err != nil {
			fmt.Println("ReadFull error", err)
			break
		}

		//拆包
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("Unpacking error", err)
			break
		}

		//是否需要二次读取
		var data []byte
		if msg.GetMsgLen() > 0 {

			//根据msg长度二次读取消息内容
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConn(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}
		msg.SetMsgData(data)

		//创建一个Request对象
		req := Request{
			conn: c,
			//请求数据
			msg: msg,
		}

		//从路由找到注册绑定的conn对应的应用
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClose {
		return errors.New("Connection is closed ")
	}

	//创建一个封包对象
	dp := NewDataPackage()

	//封装
	msg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("Packing error,msgId is=", err)
		return errors.New("pack msg error")
	}

	//发送消息
	if _, err := c.Conn.Write(msg); err != nil && err != io.EOF {
		fmt.Println("writing msgId = ", msgId, "error", err)
		return errors.New("conn write error")
	}
	return nil
}

// Start 启动链接
func (c *Connection) Start() {
	fmt.Sprintln("Coon start,ConnID=", c.ConnID)

	//TODO 启动从当前链接读业务数据
	go c.startReader()

	//TODO 启动从当前链接写业务数据
}

func (c *Connection) Stop() {

	fmt.Sprintln("Coon stop,ConnID:", c.ConnID)
	if c.isClose {
		return
	}
	c.isClose = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConn() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
