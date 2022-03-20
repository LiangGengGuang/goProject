package ziface

/*
	消息载体
*/
type IMessage interface {

	// GetMsgId 获取消息Id
	GetMsgId() uint32

	// SetMsgId 设置消息Id
	SetMsgId(uint32)

	// GetMsgLen 获取消息长度
	GetMsgLen() uint32

	// SetMsgLen 设置消息长度
	SetMsgLen(uint32)

	// GetMsgData 获取消息内容
	GetMsgData() []byte

	// SetMsgData 设置消息内容
	SetMsgData([]byte)
}
