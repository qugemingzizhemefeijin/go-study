package main

import (
	hs "../service"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RestServiceImpl struct{}

func (r *RestServiceImpl) Get(ctx context.Context, message *hs.StringMessage) (*hs.StringMessage, error) {
	return &hs.StringMessage{Value: "Get hi:" + message.Value + "#"}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *hs.StringMessage) (*hs.StringMessage, error) {
	return &hs.StringMessage{Value: "Post hi:" + message.Value + "@"}, nil
}

// https://blog.csdn.net/Mr_XiMu/article/details/125000670

// gRPC-Gateway使用指南 https://www.liwenzhou.com/posts/Go/grpc-gateway/
func main() {
	grpcServer := grpc.NewServer()
	hs.RegisterRestServiceServer(grpcServer, new(RestServiceImpl))
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	_ = grpcServer.Serve(listener)
}