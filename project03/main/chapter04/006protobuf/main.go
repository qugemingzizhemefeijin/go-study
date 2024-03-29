package main

import (
	"log"
	"net"
	"net/rpc"

	"./pb"
)

type HelloService struct {

}

func (p *HelloService) Hello(request pb.String, reply *pb.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

// protoc --go_out=. hello.proto
func main() {
	_ = rpc.Register(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}