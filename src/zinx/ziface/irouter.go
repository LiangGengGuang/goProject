package ziface

/*
	路由里的数据都是IRequest
*/
type IRouter interface {

	// PreHandle 处理conn业务之前的钩子方法
	PreHandle(IRequest)

	// Handle 处理conn业务方法
	Handle(IRequest)

	// PostHandle 处理conn业务之后的钩子方法
	PostHandle(IRequest)
}
