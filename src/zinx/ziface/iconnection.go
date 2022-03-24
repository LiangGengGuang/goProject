package ziface

import "net"

/*
	IConnection 链接载体
*/
type IConnection interface {

	// Start 启动
	Start()

	// Stop 关闭
	Stop()

	// GetTCPConn 获取当前链路的conn对象
	GetTCPConn() *net.TCPConn

	// GetConnID 获取链路ID
	GetConnID() uint32

	// RemoteAddr 获取客户端链接的地址和端口
	RemoteAddr() net.Addr

	// SendMsg 发送消息
	SendMsg(uint32, []byte) error
}

// HandleFunc 链接所绑定的处理业务的函数类型
type HandleFunc func(*net.TCPConn, []byte, int) error
