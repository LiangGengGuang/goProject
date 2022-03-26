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

// PreHandle 处理conn业务之前的钩子方法
//func (r *PingRouter) PreHandle(req ziface.IRequest) {
//
//	if _, err := req.GetConnection().GetTCPConn().Write([]byte("PreHandle...\n")); err != nil {
//		fmt.Println("PreHandle error:", err)
//	}
//}

// Handle 处理conn业务方法
func (r *PingRouter) Handle(req ziface.IRequest) {

	fmt.Println("recv from client: msgId=", req.GetMsgId(), "msg=", string(req.GetData()))

	if err := req.GetConnection().SendMsg(200, []byte("Ping...Ping...Ping...")); err != nil {
		fmt.Println("sendMsg error:", err)
	}
}

// PostHandle 处理conn业务之后的钩子方法
//func (r *PingRouter) PostHandle(req ziface.IRequest) {
//
//	if _, err := req.GetConnection().GetTCPConn().Write([]byte("PostHandle...\n")); err != nil {
//		fmt.Println("PostHandle error:", err)
//	}
//}

type CustomRouter struct {
	znet.BaseRouter
}

// Handle 自定义业务方法
func (c *CustomRouter) Handle(req ziface.IRequest) {

	fmt.Println("recv from client: msgId=", req.GetMsgId(), "msg=", string(req.GetData()))

	if err := req.GetConnection().SendMsg(201, []byte("welcome to used zinx App")); err != nil {
		fmt.Println("sendMsg error:", err)
	}
}

func DoOnConnStart(conn ziface.IConnection) {

	fmt.Println("============>DoOnConnStart")
	if err := conn.SendMsg(202, []byte("DoOnConnStart successfully...")); err != nil {
		fmt.Println(err)
	}

}
func DoOnConnStop(conn ziface.IConnection) {

	fmt.Println("============>DoOnConnStop")
	fmt.Println("connID", conn.GetConnID())
}

func main() {

	server := znet.NewServer()

	//添加router
	server.AddMsgHandler(0, &PingRouter{})
	//添加router
	server.AddMsgHandler(1, &CustomRouter{})

	//链接建立
	server.SetOnConnStart(DoOnConnStart)

	//链接销毁
	server.SetOnConnStop(DoOnConnStop)

	//运行服务器
	server.Run()
}
