package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

type Client struct {
	Name       string
	coon       net.Conn
	IPVersion  string
	ServerIp   string
	ServerPort int
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "h", "", "链接服务器ip")
	flag.IntVar(&serverPort, "p", 80, "链接服务器port")

}

func main() {
	//命令行解析
	flag.Parse()

	coon, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("dial TCP addr err:", err)
		return
	}
	for {
		dp := znet.NewDataPackage()

		//消息一
		pack1, err := dp.Pack(znet.NewMessage(0, []byte("ping zinx APP")))
		if err != nil {
			fmt.Println("Pack error", err)
			break
		}

		//消息二
		pack2, err := dp.Pack(znet.NewMessage(1, []byte("used zinx APP")))
		if err != nil {
			fmt.Println("Pack error", err)
			break
		}

		//模拟TCP粘包
		pack := append(pack1, pack2...)
		if _, err := coon.Write(pack); err != nil {
			fmt.Println("write err:", err)
			break
		}

		headLen := make([]byte, dp.GetHeadLength())
		if _, err := io.ReadFull(coon, headLen); err != nil && err != io.EOF {
			fmt.Println("ReadFull error", err)
			break
		}

		msgHead, err := dp.Unpack(headLen)
		if err != nil {
			fmt.Println("Pack error", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			//根据msg长度二次读取消息内容
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(coon, msg.Data); err != nil && err != io.EOF {
				fmt.Println("ReadFull error", err)
				break
			}
			fmt.Println("===============server call back===============")
			fmt.Printf("msgId=%d\n", msg.GetMsgId())
			fmt.Printf("msgData=%s\n", msg.GetMsgData())
		}
		time.Sleep(1 * time.Second)
	}
}
