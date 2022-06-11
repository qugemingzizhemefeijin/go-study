package main

import (
	"../pb"
	"context"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"google.golang.org/grpc"
	"time"
)


const (
	address     = "localhost:8012"
)


func main() {
	fmt.Println("client start")
	// 创建一个 gRPC 连接，用来跟服务端进行通信
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("fail to dial, error:", err)
	}
	defer conn.Close()

	// 创建 客户端对象
	c := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.Hello(ctx, &pb.String{Value: "Hello"})
	if err != nil {
		log.Fatal("say Hello error:", err)
	}

	fmt.Println("server response : ", response)
}
