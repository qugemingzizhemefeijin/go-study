package main

import (
	hs "../helloservice"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

// 服务端或客户端的单向流是双向流的特例，我们在 HelloService 增加一个支持双向流的 Channel 方法
// 关键字 stream 指定启用流特性，参数部分是接收客户端参数的流，返回值是返回给客户端的流。

/*
重新生成代码可以看到接口中新增加的 Channel 方法的定义：
type HelloServiceServer interface {
	Hello(context.Context, *String) (*String, error)
	Channel(HelloService_ChannelServer) error
}
type HelloServiceClient interface {
	Hello(ctx context.Context, in *String, opts ...grpc.CallOption) (
		*String, error,
	)
	Channel(ctx context.Context, opts ...grpc.CallOption) (
		HelloService_ChannelClient, error,
	)
}

在服务端的 Channel 方法参数是一个新的 HelloService_ChannelServer 类型的参数，可以用于和客户端双向通信。
客户端的 Channel 方法返回一个 HelloService_ChannelClient 类型的返回值，可以用于和服务端进行双向通信。

type HelloService_ChannelServer interface {
	Send(*String) error
	Recv() (*String, error)
	grpc.ServerStream
}

type HelloService_ChannelClient interface {
	Send(*String) error
	Recv() (*String, error)
	grpc.ClientStream
}

可以发现服务端和客户端的流辅助接口均定义了 Send 和 Recv 方法用于流数据的双向通信。
 */

type HelloServiceImpl struct{

}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *hs.String) (*hs.String, error) {
	reply := &hs.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

// 服务端在循环中接收客户端发来的数据，如果遇到 io.EOF 表示客户端流被关闭，如果函数退出表示服务端流关闭。
// 生成返回的数据通过流发送给客户端，双向流数据的发送和接收都是完全独立的行为。需要注意的是，发送和接收的操作并不需要一一对应，
// 用户可以根据真实场景进行组织代码。
//grpc stream
func (p *HelloServiceImpl) Channel(stream hs.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &hs.String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()
	hs.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	_ = grpcServer.Serve(listener)
}
