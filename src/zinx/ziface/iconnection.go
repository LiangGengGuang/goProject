package ziface

import "net"

type IConnection interface {

	//启动
	Start()

	//关闭
	Stop()

	//获取当前链路的conn对象
	GetTCPConn() *net.TCPConn

	//获取链路ID
	GetConnID() uint32

	//获取客户端链接的地址和端口
	RemoteAddr() net.Addr

	//发送数据
	Send(data []byte) error
}

//链接所绑定的处理业务的函数类型
type HandleFunc func(*net.TCPConn, []byte, int) error
