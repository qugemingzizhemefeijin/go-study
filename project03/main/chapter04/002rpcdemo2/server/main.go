package main

import (
	service "../service"
	"log"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct {

}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request + ", " + time.Now().String()
	return nil
}

// 比第一个案例抽象出一个公共的接口组件，接口暴露地址也作为公共变量。

// 标准库的 RPC 默认采用 Go 语言特有的 gob 编码，因此从其它语言调用 Go 语言实现的 RPC 服务将比较困难。
func main() {
	_ = service.RegisterHelloService(new(HelloService))

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
