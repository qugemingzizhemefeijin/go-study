package main

import (
	pb "../pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

/*
报错的话需要在openssl req 前面添加上 MSYS_NO_PATHCONV=1

openssl genrsa -out server.key 2048
MSYS_NO_PATHCONV=1 openssl req -new -x509 -days 3650 -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" -key server.key -out server.crt

openssl genrsa -out client.key 2048
MSYS_NO_PATHCONV=1 openssl req -new -x509 -days 3650 -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" -key client.key -out client.crt

subj 参数中的 /CN=server.grpc.io 表示服务器的名字为 server.grpc.io，在验证服务器的证书时需要用到该信息
*/

var (
	port = "5000"

	// 不知道为什么这边非要写全路径
	tlsDir = "E:/gocode/crs/go_code/project03/main/chapter04/015tls1/tls"

	server_crt = tlsDir + "/server.crt"
	server_key = tlsDir + "/server.key"
)

type myGrpcServer struct {
}

func (s *myGrpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// credentials.NewServerTLSFromFile 函数是从文件为服务器构造证书对象，然后通过
	// grpc.Creds(creds)函数将证书包装为选项后作为参数传入 grpc.NewServer 函数
	creds, err := credentials.NewServerTLSFromFile(server_crt, server_key)
	if err != nil {
		log.Panicf("could not load server key pair: %s", err)
	}

	server := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(server, new(myGrpcServer))

	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}
