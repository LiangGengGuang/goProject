package znet

import "zinx/ziface"

/*
	数据传送的消息载体对象
*/
type Message struct {

	//id
	MsgId uint32
	//消息长度
	MsgLen uint32
	//消息内容
	data []byte
}

// NewMessage 初始化消息对象
func NewMessage(msgId uint32, data []byte) ziface.IMessage {

	return &Message{
		MsgId:  msgId,
		MsgLen: uint32(len(data)),
		data:   data,
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.MsgId
}

func (m *Message) SetMsgId(id uint32) {
	m.MsgId = id
}

func (m *Message) GetMsgLen() uint32 {
	return m.MsgLen
}

func (m *Message) SetMsgLen(msgLen uint32) {
	m.MsgLen = msgLen
}

func (m *Message) GetMsgData() []byte {
	return m.data
}

func (m *Message) SetMsgData(data []byte) {
	m.data = data
}
