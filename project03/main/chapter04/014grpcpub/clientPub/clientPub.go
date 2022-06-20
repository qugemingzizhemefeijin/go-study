package main

import (
	pb "../pubsubservice"
	"context"
	"google.golang.org/grpc"
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

	// 客户端发布信息 golang :hello Go
	_, err = client.Publish(context.Background(), &pb.String{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}

	// 客户端发布信息 docker: hello Docker
	_, err = client.Publish(context.Background(), &pb.String{Value: "docker: hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}