package main

import (
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	rpc2 "grpc/client/rpc"
	"log"
	"time"
)

// @Description
// @Author lianggengguang
// @Date 2023/3/5

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name = flag.String("name", "Leung", "Name to greet")
)

func main() {
	flag.Parse()
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc2.NewGreeterClient(conn)

	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &rpc2.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting receive : %s", r.GetReply())
}
