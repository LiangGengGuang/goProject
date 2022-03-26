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
	Connection 链接
*/
type Connection struct {
	//当前Conn属于哪个Server
	tcpServer ziface.IServer
	//当前链接的socket
	conn *net.TCPConn
	//当前链接ID
	connID uint32
	//当前链接是否关闭
	isClose bool
	//告知当前链接已经停止/退出
	exitChan chan bool
	//无缓冲，用于读写之间的消息通信
	msgChan chan []byte
	//该链接消息管理
	msgHandler ziface.IMsgHandler
}

// NewConnection 初始化链路模块方法
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) ziface.IConnection {

	c := &Connection{
		tcpServer:  server,
		conn:       conn,
		connID:     connID,
		isClose:    false,
		exitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
		msgHandler: msgHandler,
	}

	//将链接添加进容器
	c.tcpServer.GetConnMgr().Add(c)
	return c
}

//读取链接请求数据
func (c *Connection) startReader() {

	fmt.Println("[reader goroutine is running]")
	defer fmt.Println("reader goroutine exit,ConnID:", c.connID, "RemoteAddr", c.RemoteAddr().String())
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
			c.msgHandler.SendMsgToTaskQueue(&req)
		} else {
			//从路由找到注册绑定的conn对应的应用
			c.msgHandler.DoMsgHandler(&req)
		}
	}
}

//向链接请求写入数据发送给客户端
func (c *Connection) startWriter() {

	fmt.Println("[writer goroutine is running]")

	defer fmt.Println("writer goroutine exit,ConnID:", c.connID, "RemoteAddr", c.RemoteAddr().String())

	for {
		select {
		case data, ok := <-c.msgChan:
			if ok {
				if _, err := c.conn.Write(data); err != nil {
					fmt.Println("sendMsg write error:", err)
					break
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}

		case <-c.exitChan:
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
	fmt.Sprintln("Coon start,ConnID=", c.connID)

	//启动从当前链接读业务数据
	go c.startReader()

	//启动从当前链接写业务数据
	go c.startWriter()

	//按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	c.tcpServer.CallOnConnStart(c)
}

func (c *Connection) Stop() {

	fmt.Println("Coon stop,ConnID:", c.connID)

	//如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用
	c.tcpServer.CallOnConnStop(c)

	//回收资源
	defer close(c.exitChan)
	defer close(c.msgChan)

	if c.isClose {
		return
	}
	c.isClose = true

	//关闭socket链接
	_ = c.conn.Close()

	//告知Writer关闭
	c.exitChan <- true

	//将链接从容器中移除
	c.tcpServer.GetConnMgr().Remove(c.connID)
}

func (c *Connection) GetTCPConn() *net.TCPConn {
	return c.conn
}

func (c *Connection) GetConnID() uint32 {
	return c.connID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
