package znet

import "zinx/ziface"

/*
	实现router时，先嵌入BaseRouter基类，然后根据要求对整个基类进行方法的重写
*/
type BaseRouter struct {
}

// PreHandle 处理conn业务之前的钩子方法
func (r BaseRouter) PreHandle(_ ziface.IRequest) {
}

// Handle 处理conn业务方法
func (r BaseRouter) Handle(_ ziface.IRequest) {
}

// PostHandle 处理conn业务之后的钩子方法
func (r BaseRouter) PostHandle(_ ziface.IRequest) {
}
