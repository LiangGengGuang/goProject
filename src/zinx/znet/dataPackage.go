package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/utils"
	"zinx/ziface"
)

/*
	DataPackage 解决TCP数据粘包的处理对象
*/
type DataPackage struct{}

func NewDataPackage() *DataPackage {

	return &DataPackage{}
}

// HeadLength uint32(4字节) + MsgId uint32(4字节)
const HeadLength = 8

func (dp *DataPackage) GetHeadLength() uint32 {

	return HeadLength
}

func (dp *DataPackage) Pack(msg ziface.IMessage) ([]byte, error) {

	//创建一个写入byte[]字节的缓冲
	buffer := bytes.NewBuffer([]byte{})

	//1.先写入消息的长度
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	//2.在写入消息的ID
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//3.最后写入消息的内容
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (dp *DataPackage) Unpack(binaryData []byte) (ziface.IMessage, error) {

	//创建一个读取二进制数据的reader
	buffer := bytes.NewReader(binaryData)

	msg := &Message{}

	//1.先读取消息的长度
	if err := binary.Read(buffer, binary.LittleEndian, &msg.MsgLen); err != nil {
		return nil, err
	}

	//2.在读取消息的ID
	if err := binary.Read(buffer, binary.LittleEndian, &msg.MsgId); err != nil {
		return nil, err
	}

	//判断消息的长度是否超过配置的最大长度限制
	maxPackageSize := utils.GlobalObject.MaxPackageSize
	if maxPackageSize > 0 && maxPackageSize < msg.MsgLen {
		return nil, errors.New("too Large msg data receive")
	}
	return msg, nil
}
