package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/utils"
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
	//无缓冲，用于读写之间的消息通信
	msgChan chan []byte
	//该链接消息管理
	MsgHandler ziface.IMsgHandler
}

// NewConnection 初始化链路模块方法
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) ziface.IConnection {

	return &Connection{
		Conn:       conn,
		ConnID:     connID,
		isClose:    false,
		MsgHandler: msgHandler,
		msgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
	}
}

//读取链接请求数据
func (c *Connection) startReader() {

	fmt.Println("[reader goroutine is running]")
	defer fmt.Println("reader goroutine exit,ConnID:", c.ConnID, "RemoteAddr", c.RemoteAddr().String())
	defer c.Stop()

	for {

		//创建一个拆包对象
		dp := NewDataPackage()

		//读取id和消息长度
		headData := make([]byte, dp.GetHeadLength())
		if _, err := io.ReadFull(c.GetTCPConn(), headData); err != nil {
			fmt.Println("read msg head error:", err)
			break
		}

		//拆包
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("Unpacking error:", err)
			break
		}

		//是否需要二次读取
		var data []byte
		if msg.GetMsgLen() > 0 {

			//根据msg长度二次读取消息内容
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConn(), data); err != nil {
				fmt.Println("read msg data error:", err)
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

		if utils.GlobalObject.WorkerPoolSize > 0 {
			//已经启动工作池机制，将消息交给Worker处理
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//从路由找到注册绑定的conn对应的应用
			c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

//向链接请求写入数据发送给客户端
func (c *Connection) startWriter() {

	fmt.Println("[writer goroutine is running]")

	defer fmt.Println("writer goroutine exit,ConnID:", c.ConnID, "RemoteAddr", c.RemoteAddr().String())

	for {
		select {
		case data, ok := <-c.msgChan:
			if ok {
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("sendMsg write error:", err)
					break
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}

		case <-c.ExitChan:
			//代表Reader已经退出，Reader一并退出
			return
		}
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
		fmt.Println("msgId is=", msgId, "Packing error:", err)
		return errors.New("pack msg error")
	}

	//消息发给管道
	c.msgChan <- msg
	return nil
}

// Start 启动链接
func (c *Connection) Start() {
	fmt.Sprintln("Coon start,ConnID=", c.ConnID)

	//TODO 启动从当前链接读业务数据
	go c.startReader()

	//TODO 启动从当前链接写业务数据
	go c.startWriter()
}

func (c *Connection) Stop() {

	fmt.Println("Coon stop,ConnID:", c.ConnID)

	if c.isClose {
		return
	}
	c.isClose = true

	//关闭socket链接
	c.Conn.Close()

	//告知Writer关闭
	c.ExitChan <- true

	//回收资源
	close(c.ExitChan)
	close(c.msgChan)
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
