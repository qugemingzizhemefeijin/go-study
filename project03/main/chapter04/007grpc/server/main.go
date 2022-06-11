package main

import (
	"../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (p *HelloService) Hello(ctx context.Context, request *pb.String) (response *pb.String, err error) {
	log.Printf("接收到客户端信息: %v\n", request.GetValue())
	response = &pb.String{
		Value: "hello:" + request.GetValue(),
	}
	return response, nil
}

// protoc --go_out=plugins=grpc:. hello.proto
func main() {
	fmt.Println("start server")
	// 监听客户端的请求
	listen, err := net.Listen("tcp", ":8012")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	// 创建一个 gRPC Server 实例
	server := grpc.NewServer()

	// 将该服务注册到 gRPC 框架中
	pb.RegisterHelloServiceServer(server, &HelloService{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	// 启动 gRPC 服务
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("Serve error:", err)
	}
}
