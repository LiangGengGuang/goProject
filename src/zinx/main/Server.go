package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

/*
	基于zinx框架开发的 服务器应用程序
*/

type PingRouter struct {
	znet.BaseRouter
}

//处理conn业务之前的钩子方法
func (this *PingRouter) PreHandle(req ziface.IRequest) {

	if _, err := req.GetConnection().GetTCPConn().Write([]byte("PreHandle...\n")); err != nil {
		fmt.Println("PreHandle error:", err)
	}
}

//处理conn业务方法
func (this *PingRouter) Handle(req ziface.IRequest) {

	if _, err := req.GetConnection().GetTCPConn().Write([]byte("Handle...\n")); err != nil {
		fmt.Println("Handle error:", err)
	}
}

//处理conn业务之后的钩子方法
func (this *PingRouter) PostHandle(req ziface.IRequest) {

	if _, err := req.GetConnection().GetTCPConn().Write([]byte("PostHandle...\n")); err != nil {
		fmt.Println("PostHandle error:", err)
	}
}

func main() {

	server := znet.NewServer()
	//添加router
	server.AddRouter(&PingRouter{})
	//运行服务器
	server.Run()
}
