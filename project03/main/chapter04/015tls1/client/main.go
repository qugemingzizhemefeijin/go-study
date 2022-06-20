package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
)

var (
	port = "5000"

	// 不知道为什么这边非要写全路径
	tlsDir = "E:/gocode/crs/go_code/project03/main/chapter04/015tls1/tls"

	server_crt = tlsDir + "/server.crt"
)

// 1.15.1版本  ，由于需要用到GRPC 所以就开始写代码 然后就碰到x509: certificate relies on legacy Common Name field
// GO1.15   X509 被砍了（不能用了） ，需要用到SAN证书
func main() {
	// credentials.NewClientTLSFromFile 是构造客户端用的证书对象，第一个参数是服务器的证书文件，第二个参数是签发证书的服务器的名字。
	// 然后通过 grpc.WithTransportCredentials(creds) 将证书对象转为参数选项传人 grpc.Dial 函数。
	creds, err := credentials.NewClientTLSFromFile(server_crt, "server.grpc.io")
	if err != nil {
		log.Panicf("could not load client key pair: %s", err)
	}

	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "gopher"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("doClientWork: %s", r.Message)
}
