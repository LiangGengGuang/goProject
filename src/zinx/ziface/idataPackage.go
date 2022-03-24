package ziface

/*
	IDataPackage tcp的TLV格式，解决TCP数据粘包
*/
type IDataPackage interface {

	// GetHeadLength head长度
	GetHeadLength() uint32

	// Pack 封装
	Pack(IMessage) ([]byte, error)

	// Unpack 拆包
	Unpack([]byte) (IMessage, error)
}
