package ziface

/*
	服务器接口
*/
type IServer interface {

	// Start 启动
	Start()

	// Stop 停止
	Stop()

	// Run 运行
	Run()

	// AddMsgHandler 给当前服务注册一个路由方法，共客户端链接使用
	AddMsgHandler(uint32, IRouter)

	// GetConnMgr 获取链接容器
	GetConnMgr() IConnManager
}
