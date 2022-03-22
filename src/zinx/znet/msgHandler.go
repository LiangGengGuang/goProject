package znet

import (
	"fmt"
	"strconv"
	"zinx/utils"
	"zinx/ziface"
)

/*
	消息管理具体实现
*/
type MsgHandler struct {
	//存放每个msgId处理的方法
	Apis map[uint32]ziface.IRouter
	//消息队列
	TaskQueue []chan ziface.IRequest
	//工作池的大小
	WorkPoolSize uint32
}

func NewMsgHandler() ziface.IMsgHandler {

	return &MsgHandler{
		Apis:         make(map[uint32]ziface.IRouter),
		TaskQueue:    make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkPoolSize: utils.GlobalObject.WorkerPoolSize,
	}
}

func (m *MsgHandler) DoMsgHandler(request ziface.IRequest) {

	router, ok := m.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgId(), " is not FOUND!")
		return
	}

	//执行对应处理方法
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

func (m *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {

	if _, ok := m.Apis[msgId]; ok {
		panic("repeated api , msgID = " + strconv.Itoa(int(msgId)))
	}
	m.Apis[msgId] = router
}

// SendMsgToTaskQueue 启动一个工作流程
func (m *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {

	//根据ConnID来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配法则,得到需要处理此条连接的workerID
	workerID := request.GetConnection().GetConnID() % m.WorkPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(), " request msgID=", request.GetMsgId(), "to workerID=", workerID)
	//将请求消息发送给任务队列
	m.TaskQueue[workerID] <- request
}

// StartWorkPool 启动一个工作池（开启工作池只在系统启动时发生一次）
func (m *MsgHandler) StartWorkPool() {

	for i := 0; i < int(m.WorkPoolSize); i++ {

		//一个work被启动，开辟work存放的消息队列的数量
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerPoolSize)
		go m.startOneWorkPool(i, m.TaskQueue[i])
	}

}

//启动一个工作流程
func (m *MsgHandler) startOneWorkPool(workId int, taskQueue chan ziface.IRequest) {

	fmt.Println("workId=", workId, "is started...")
	for {
		select {
		//如果有消息过来，执行当前request所绑定业务
		case request, ok := <-taskQueue:
			if ok {
				m.DoMsgHandler(request)
			} else {
				fmt.Println("taskQueue is error")
			}
		}
	}
}
