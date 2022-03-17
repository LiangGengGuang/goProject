package ziface

type IRequest interface {

	//获取请求链接
	GetConnection() IConnection

	//获取请求数据
	GetData() []byte
}
