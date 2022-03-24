package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

var serverIp2 string
var serverPort2 int

func init() {
	flag.StringVar(&serverIp2, "h", "0.0.0.0", "链接服务器ip")
	flag.IntVar(&serverPort2, "p", 8989, "链接服务器port")

}

func main() {
	//命令行解析
	flag.Parse()

	coon, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp2, serverPort2))
	if err != nil {
		fmt.Println("dial TCP addr err:", err)
		return
	}
	for {
		dp := znet.NewDataPackage()

		//消息
		pack, err := dp.Pack(znet.NewMessage(1, []byte("used zinx APP")))
		if err != nil {
			fmt.Println("Pack error:", err)
			break
		}
		if _, err := coon.Write(pack); err != nil {
			fmt.Println("write err:", err)
			break
		}

		headLen := make([]byte, dp.GetHeadLength())
		if _, err := io.ReadFull(coon, headLen); err != nil {
			fmt.Println("client2 ReadFull error:", err)
			break
		}

		msgHead, err := dp.Unpack(headLen)
		if err != nil {
			fmt.Println("Pack error:", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			//根据msg长度二次读取消息内容
			msg := msgHead.(*znet.Message)
			msg.SetMsgData(make([]byte, msg.GetMsgLen()))
			if _, err := io.ReadFull(coon, msg.GetMsgData()); err != nil {
				fmt.Println("client2 ReadFull error:", err)
				break
			}
			fmt.Println("===============server call back===============")
			fmt.Printf("msgId=%d\n", msg.GetMsgId())
			fmt.Printf("msgData=%s\n", msg.GetMsgData())
		}
		time.Sleep(1 * time.Second)
	}
}
