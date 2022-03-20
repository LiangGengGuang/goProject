package znet

import "zinx/ziface"

/*
	链接请求实体对象
*/
type Request struct {
	//链接
	conn ziface.IConnection

	//请求数据
	msg ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {

	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
