package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

/*
	服务器实体对象
*/
type Server struct {
	Name       string
	TcpVersion string
	IP         string
	Port       int
	MsgHandler ziface.IMsgHandler
}

// NewServer 初始化server模块
func NewServer() ziface.IServer {
	return &Server{
		Name:       utils.GlobalObject.Name,
		TcpVersion: "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandler(),
	}
}

// TCPListener 处理tcp监听器内容
func (s *Server) TCPListener(listener net.TCPListener) {
	var connID uint32 = 0

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("accept TCP err：", err)
			continue
		}

		//将处理新链接的业务与conn进行绑定
		connection := NewConnection(conn, connID, s.MsgHandler)
		connID++

		//启动当前链接的业务处理
		go connection.Start()
	}
}

// Start 服务器接口启动实现方法
func (s *Server) Start() {
	//TODO 启动server服务,解析TCP请求并进行监听
	fmt.Printf("[start] Server Name %s,Listener at IP：%s,Port：%d\n", s.Name, s.IP, s.Port)

	//获取tcp的address
	go func() {
		addr, err := net.ResolveTCPAddr(s.TcpVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve TCP addr err：", err)
			return
		}

		//监听服务器地址
		listener, err := net.ListenTCP(s.TcpVersion, addr)
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

	//TODO 将服务器的资源、状态、开辟的链接信息，进行停止或回收

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
	s.MsgHandler.AddRouter(msgID, router)
}
