package main

// go get -t github.com/docker/docker/pkg/pubsub

import (
	pb "../pubsubservice"
	"context"
	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	// 构造一个发布对象
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

// Publish 实现发布方法
func (p *PubsubService) Publish(ctx context.Context, arg *pb.String) (*pb.String, error) {
	// 发布消息
	p.pub.Publish(arg.GetValue())
	// debug
	// reply := &String{Value: "<Publish>  " + arg.GetValue()}
	// fmt.Println(reply.GetValue())
	return &pb.String{}, nil
}

// Subscribe 实现订阅方法
func (p *PubsubService) Subscribe(arg *pb.String, stream pb.PubsubService_SubscribeServer) error {
	// SubscribeTopic 增加一个使用函数过滤器的订阅者
	// func(v interface{}) 定义函数过滤的规则
	// SubscribeTopic 返回一个chan interface{}
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		// 接收数据是string，并且key是以arg为前缀的
		if key, ok := v.(string); ok {
			// debug
			// fmt.Printf("<debug> %t %s %s %t\n",
			// ok,arg.GetValue(),key,strings.HasPrefix(key,arg.GetValue()))
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	// 服务器遍历chan，并将其中信息发送给订阅客户端
	for v := range ch {
		if err := stream.Send(&pb.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterPubsubServiceServer(grpcServer, NewPubsubService())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	_ = grpcServer.Serve(lis)
}