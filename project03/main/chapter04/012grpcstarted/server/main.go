package main

import (
	hs "../helloservice"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 基于服务端的 HelloServiceServer 接口可以重新实现 HelloService 服务
type HelloServiceImpl struct {

}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *hs.String) (*hs.String, error) {
	reply := &hs.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

// gRPC 是 Google 公司基于 Protobuf 开发的跨语言的开源 RPC 框架。gRPC 基于 HTTP/2 协议设计，
// 可以基于一个 HTTP/2 链接提供多个服务，对于移动设备更加友好。本节将讲述 gRPC 的简单用法。

// 最底层为 TCP 或 Unix Socket 协议，在此之上是 HTTP/2 协议的实现，然后在 HTTP/2 协议之上又构建了针对 Go 语言的 gRPC 核心库。
// 应用程序通过 gRPC 插件生产的 Stub 代码和 gRPC 核心库通信，也可以直接和 gRPC 核心库通信。

// gRPC 和标准库的 RPC 框架有一个区别，gRPC 生成的接口并不支持异步调用。
// 不过我们可以在多个 Goroutine 之间安全地共享 gRPC 底层的 HTTP/2 链接，因此可以通过在另一个 Goroutine 阻塞调用的方式模拟异步调用。
func main() {
	// 构造一个 gRPC 服务对象
	grpcServer := grpc.NewServer()
	// 通过 gRPC 插件生成的 RegisterHelloServiceServer 函数注册我们实现的 HelloServiceImpl 服务
	hs.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// 监听端口上提供 gRPC 服务
	_ = grpcServer.Serve(listener)
}
