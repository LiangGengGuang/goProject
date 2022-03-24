package ziface

/*
	IMsgHandler 消息管理抽象层
*/
type IMsgHandler interface {

	// DoMsgHandler 调度/执行对应router消息处理方法
	DoMsgHandler(IRequest)

	// AddRouter 为消息添加具体的处理逻辑
	AddRouter(uint32, IRouter)

	// StartWorkPool 开启工作池
	StartWorkPool()

	// SendMsgToTaskQueue 将请求消息传递给工作流
	SendMsgToTaskQueue(IRequest)
}
