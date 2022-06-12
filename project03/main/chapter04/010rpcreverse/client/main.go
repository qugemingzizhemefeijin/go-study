package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net"
	"net/rpc"
)

// 反向RPC
func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error:", err)
			}

			// 这里有点像是将服务端的conn包装成一个client
			clientChan <- rpc.NewClient(conn)
		}
	}()

	client := <- clientChan
	defer client.Close()

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal("", err)
	}

	fmt.Println(reply)
}
