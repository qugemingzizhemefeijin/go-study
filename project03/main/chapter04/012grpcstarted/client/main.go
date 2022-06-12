package main

import (
	hs "../helloservice"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

// grpc入门案例
func main() {
	// 负责和 gRPC 服务建立链接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 基于已经建立的链接构造 HelloServiceClient 对象
	client := hs.NewHelloServiceClient(conn)
	// 返回的 client 其实是一个 HelloServiceClient 接口对象，通过接口定义的方法就可以调用服务端对应的 gRPC 服务提供的方法。
	reply, err := client.Hello(context.Background(), &hs.String{
		Value: "Hello",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())
}