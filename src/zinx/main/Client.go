package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"
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
		if _, err := coon.Write([]byte("client link request\n")); err != nil {
			fmt.Println("write err:", err)
			return
		}

		buff := make([]byte, 512)
		read, err := coon.Read(buff)
		if err != nil && err != io.EOF {
			fmt.Println("write err:", err)
			return
		}
		fmt.Println("===============server call back")
		fmt.Printf("%s", buff)
		fmt.Println("read=", read)

		time.Sleep(1 * time.Second)
	}
}
