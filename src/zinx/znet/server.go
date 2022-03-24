package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

/*
	Server 服务器实体对象
*/
type Server struct {
	name       string
	tcpVersion string
	ip         string
	port       int
	msgHandler ziface.IMsgHandler
	connMgr    ziface.IConnManager
}

// NewServer 初始化server模块
func NewServer() ziface.IServer {
	return &Server{
		name:       utils.GlobalObject.Name,
		tcpVersion: "tcp4",
		ip:         utils.GlobalObject.Host,
		port:       utils.GlobalObject.TcpPort,
		msgHandler: NewMsgHandler(),
		connMgr:    NewConnManager(),
	}
}

// TCPListener 处理tcp监听器内容
func (s *Server) TCPListener(listener net.TCPListener) {
	var connID uint32 = 0

	for {
		TCPConn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("accept TCP err：", err)
			continue
		}

		//判断是否超链接最大支持数
		if s.connMgr.Quantity() >= utils.GlobalObject.MaxCon {

			_ = TCPConn.Close()
			fmt.Println("=======>the maximum number of links has been exceeded,max conn=", utils.GlobalObject.MaxCon)
			//TODO 给客户端发送超出最大链接数的错误包

			continue
		}

		//将处理新链接的业务与conn进行绑定
		conn := NewConnection(s, TCPConn, connID, s.msgHandler)

		//将链接添加进容器
		s.connMgr.Add(conn)
		fmt.Println("link quantity = ", s.connMgr.Quantity())

		//启动当前链接的业务处理
		go conn.Start()
	}
}

// Start 服务器接口启动实现方法
func (s *Server) Start() {
	//TODO 启动server服务,解析TCP请求并进行监听
	fmt.Printf("[START] server name %s,Listener at IP：%s,Port：%d\n", s.name, s.ip, s.port)

	//获取tcp的address
	go func() {

		//TODO 启动工作池
		s.msgHandler.StartWorkPool()

		addr, err := net.ResolveTCPAddr(s.tcpVersion, fmt.Sprintf("%s:%d", s.ip, s.port))
		if err != nil {
			fmt.Println("resolve TCP addr err：", err)
			return
		}

		//监听服务器地址
		listener, err := net.ListenTCP(s.tcpVersion, addr)
		if err != nil {
			fmt.Println("listen TCP err：", err)
			return
		}
		fmt.Println("start Zinx server successfully, Listening...")

		// 阻塞等待客户端链接，处理客户端业务
		go s.TCPListener(*listener)
	}()
}

// Stop 服务器接口停止实现方法
func (s *Server) Stop() {

	fmt.Println("[STOP] Zinx server name=", s.name)
	//TODO 将服务器的资源、状态、开辟的链接信息，进行停止或回收
	s.connMgr.ClearAll()
}

// Run 服务器接口运行实现方法
func (s *Server) Run() {

	//启动Server服务
	s.Start()

	//TODO 其他业务

	//阻塞
	select {}
}

// AddMsgHandler 给当前服务注册一个路由方法到消息管理中，供客户端链接使用
func (s *Server) AddMsgHandler(msgID uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgID, router)
}

// GetConnMgr 获取链接容器
func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.connMgr
}
