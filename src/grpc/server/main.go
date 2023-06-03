package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	rpc2 "grpc/client/rpc"
	"net"
)

// hello server

type server struct {
	rpc2.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *rpc2.HelloRequest) (*rpc2.HelloResponse, error) {
	return &rpc2.HelloResponse{Reply: "Hello " + in.Name}, nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                    // 创建gRPC服务器
	rpc2.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
	// 启动服务
	if s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
	fmt.Printf("server staring...")
}
