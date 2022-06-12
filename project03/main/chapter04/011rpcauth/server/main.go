package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
	conn    net.Conn
	isLogin bool
}

func ServeHelloService(conn net.Conn) {
	p := rpc.NewServer()
	// 向rpc server对象注册一个HelloService对象，注册后，client就可以调用HelloService的方法
	_ = p.Register(&HelloService{conn: conn})
	p.ServeConn(conn)
}

func (p *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		log.Println("login failed")
		return fmt.Errorf("auth failed")
	}

	log.Println("login ok")
	p.isLogin = true
	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
	if !p.isLogin {
		return fmt.Errorf("please login")
	}

	*reply = "hello:" + request + ", from " + p.conn.RemoteAddr().String()
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go ServeHelloService(conn)
	}
}
