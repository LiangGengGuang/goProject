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

	// SetProperty 设置链接属性
	SetProperty(string, interface{})

	// GetProperty 获取链接属性
	GetProperty(string) (interface{}, error)

	// RemoveProperty 移除链接属性
	RemoveProperty(string)
}