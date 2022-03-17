package znet

import "zinx/ziface"

type Request struct {
	//链接
	conn ziface.IConnection

	//请求数据
	data []byte
}

//初始化请求链接
func NewRequest(iConn ziface.IConnection, reqData []byte) ziface.IRequest {
	r := &Request{
		conn: iConn,
		data: reqData,
	}
	return r

}
func (r *Request) GetConnection() ziface.IConnection {

	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
