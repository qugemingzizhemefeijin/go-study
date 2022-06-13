package main

import (
	hs "../HelloService"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hs.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &hs.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

	//grpc stream
	// 客户端需要先调用 Channel 方法获取返回的流对象
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 在客户端我们将发送和接收操作放到两个独立的Goroutine。首先是向服务端发送数据：
	go func() {
		for {
			if err := stream.Send(&hs.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	// 这里是服务器将流数据发送回客户端的接收代码
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
