package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var serverIp string
var serverPort int

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	coon       net.Conn
	flag       int
}

//创建一个客户端
func NewClient(ServerIp string, ServerPort int) *Client {
	coon, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ServerIp, ServerPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client := &Client{
		ServerIp:   ServerIp,
		ServerPort: ServerPort,
		flag:       9999,
		coon:       coon,
	}
	return client
}

//公聊
func (this *Client) publicChat() {

	var publicMsg string
	fmt.Println(">>>> 请输入消息内容,输入exit退出：")
	fmt.Scanln(&publicMsg)

	for publicMsg != "exit" {
		if len(publicMsg) != 0 {

			msg := publicMsg + "\n"
			_, err := this.coon.Write([]byte (msg))
			if err != nil {
				fmt.Println("消息发送失败：", err)
				break
			}
		}
		fmt.Println(">>>> 请输入消息内容,输入exit退出：")
		fmt.Scanln(&publicMsg)
	}
}

//获取在线用户
func (this *Client) queryOnlineUser() {
	_, err := this.coon.Write([]byte ( "online\n"))
	if err != nil {
		fmt.Println("获取在线用户失败：", err)
		return
	}
}

//私聊
func (this *Client) privateChat() {

	var privateObject string
	var privateMsg string

	this.queryOnlineUser()
	fmt.Println(">>>> 请选择私聊用户,输入exit退出：")
	fmt.Scanln(&privateObject)

	for privateObject != "exit" {

		fmt.Println(">>>> 请输入消息内容,输入exit退出：")
		fmt.Scanln(&privateMsg)

		for privateMsg != "exit" {
			if len(privateMsg) != 0 {

				msg := "to|" + privateObject + "|" + privateMsg + "\n"
				_, err := this.coon.Write([]byte (msg + "\n"))
				if err != nil {
					fmt.Println("消息发送失败：", err)
					break
				}
			}
			fmt.Println(">>>> 请输入消息内容,输入exit退出：")
			fmt.Scanln(&privateMsg)
		}
		this.queryOnlineUser()
		fmt.Println(">>>> 请选择私聊用户,输入exit退出：")
		fmt.Scanln(&privateObject)
	}
}

//更改用户名
func (this *Client) updateName() bool {

	var newName string
	for len(newName) == 0 {
		fmt.Println(">>>> 请输入用户名:")
		fmt.Scanln(&newName)
	}
	//消息结尾要加\n
	msg := "rename|" + newName + "\n"

	fmt.Println("msg：", msg)

	_, err := this.coon.Write([]byte(msg))
	if err != nil {
		fmt.Println("用户名更新失败：", err)
		return false
	}
	this.Name = newName
	return true
}

//获取用户信息
func (this *Client) getInfo(newName string) {

	//消息结尾要加\n
	msg := "get|" + newName + "\n"

	fmt.Println("msg：", msg)

	_, err := this.coon.Write([]byte(msg))
	if err != nil {
		fmt.Println("获取用户信息失败：", err)
		return
	}
}

//菜单
func (this *Client) menu() bool {

	var flag int
	fmt.Println("\n ========= 菜单 =========")
	fmt.Println("1.公聊")
	fmt.Println("2.私聊")
	fmt.Println("3.更改用户名")
	fmt.Println("4.查询用户信息")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag < 0 || flag > 4 {
		fmt.Println(">>>> 请输入合法数字 <<<<")
		return false
	}
	this.flag = flag
	return true
}

//接收服务端发来的消息 单独执行
func (this *Client) receiveServerMsg() {
	//一旦client.coon有消息，就直接copy都Stdout标准输出中，永久阻塞监听
	io.Copy(os.Stdout, this.coon)
}

func (this *Client) run() {

	for this.flag != 0 {
		for this.menu() != true {
		}
		switch this.flag {
		case 1:
			//公聊
			this.publicChat()
		case 2:
			//私聊
			this.privateChat()
		case 3:
			//更改用户名
			this.updateName()
		case 4:
			//获取用户信息
			this.getInfo(this.Name)
		}
	}
}

//./client -h 127.0.0.1 -p 8888
func init() {
	flag.StringVar(&serverIp, "h", "", "链接服务器ip")
	flag.IntVar(&serverPort, "p", 80, "链接服务器port")
}

func main() {
	//命令行解析
	flag.Parse()

	if serverIp == "" || serverPort == 80 {
		fmt.Println("链接服务器失败...")
		return
	}

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("链接服务器失败...")
		return
	}

	go client.receiveServerMsg()

	fmt.Println("链接服务器成功...")
	client.run()
}
