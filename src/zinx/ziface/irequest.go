package ziface

/*
	IRequest 链接请求载体
*/
type IRequest interface {

	// GetConnection 获取请求链接
	GetConnection() IConnection

	// GetData 获取请求消息数据
	GetData() []byte

	// GetMsgId 获取请求数据ID
	GetMsgId() uint32
}
