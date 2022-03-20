package znet

import (
	"fmt"
	"strconv"
	"zinx/ziface"
)

/*
	消息管理具体实现
*/
type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() ziface.IMsgHandler {

	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
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
