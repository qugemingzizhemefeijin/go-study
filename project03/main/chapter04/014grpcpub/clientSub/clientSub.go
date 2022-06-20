package main

import (
	pb "../pubsubservice"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 新建一个客户端
	client := pb.NewPubsubServiceClient(conn)

	// 订阅服务，传入参数是 golang:
	// 会通过滤器函数，订阅者应该收到的信息为 golang: hello Go
	stream, err := client.Subscribe(context.Background(), &pb.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}

	// 阻塞遍历流，输出结果
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}
}
