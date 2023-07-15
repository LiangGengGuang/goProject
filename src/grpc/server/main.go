package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/client/rpc"
	"net"
)

// hello server

type server struct {
	rpc.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *rpc.HelloRequest) (*rpc.HelloResponse, error) {
	return &rpc.HelloResponse{Reply: "Hello " + in.Name}, nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                   // 创建gRPC服务器
	rpc.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
	// 启动服务
	if s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
	fmt.Printf("server staring...")
}
