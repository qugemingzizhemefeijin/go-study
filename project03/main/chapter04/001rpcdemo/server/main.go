package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {}

// Hello⽅法必须满⾜Go语⾔的RPC规则： ⽅法只能有两个可序列化的参数， 其中第⼆个参数是指针类型， 并且返回⼀个error类型， 同时必须是公开的⽅法。
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func calcAlg(colunm int64) string {
	idx := colunm & 0x3FC00000
	return fmt.Sprintf("#%d", idx >> 22)
}

// 其中 rpc.Register 函数调用会将对象类型中所有满足 RPC 规则的对象方法注册为 RPC 函数，所有注册的方法会放在 “HelloService” 服务空间之下。
// 然后我们建立一个唯一的 TCP 链接，并且通过 rpc.ServeConn 函数在该 TCP 链接上为对方提供 RPC 服务。
func main() {
	fmt.Println(calcAlg(16455595008))

	// new(File) 和 &File{} 是等价的
	_ = rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}