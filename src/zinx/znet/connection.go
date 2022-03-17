package znet

import (
	"fmt"
	"io"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

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

//初始化链路模块方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {

	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClose:  false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}

//读取链接请求数据
func (c *Connection) startReader() {

	defer fmt.Println("connID =", c.ConnID, "reader is exit,remote is = ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		buff := make([]byte, utils.GlobalObject.MaxPackageSize)
		if _, err := c.Conn.Read(buff); err != nil && err != io.EOF {
			fmt.Println("Read is err=", err)
			continue
		}

		req := Request{
			conn: c,
			//请求数据
			data: buff,
		}
		//从路由找到注册绑定的conn对应的应用
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

//启动链接
func (c *Connection) Start() {
	fmt.Sprintln("Coon start,ConnID=", c.ConnID)

	//TODO 启动从当前链接读业务数据
	go c.startReader()

	//TODO 启动从当前链接写业务数据
}

func (c *Connection) Stop() {

	fmt.Sprintln("Coon stop,ConnID:", c.ConnID)
	if c.isClose == true {
		return
	}
	c.isClose = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConn() *net.TCPConn {
	return c.Conn
}

func (c Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c Connection) Send(data []byte) error {
	return nil
}
