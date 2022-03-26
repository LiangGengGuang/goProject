package ziface

/*
	IConnManager 链接管理抽象方法
*/
type IConnManager interface {

	// Add 添加链接
	Add(IConnection)

	// Remove 删除链接(客户端自己断开)
	Remove(uint32)

	// Get 根据ConnId获取链接
	Get(uint32) (IConnection, error)

	// Quantity 链接数量
	Quantity() int

	// ClearAll 清空链接(主动)
	ClearAll()

	// ClearOne 清空一个链接(主动)
	ClearOne(uint32)
}
