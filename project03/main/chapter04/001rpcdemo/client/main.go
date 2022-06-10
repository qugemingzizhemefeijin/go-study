package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// ⾸选是通过rpc.Dial拨号RPC服务， 然后通过client.Call调⽤具体的RPC⽅法。 在调⽤client.Call时，
// 第⼀个参数是⽤点号链接的RPC服务名字和⽅法名字， 第⼆和第三个参数分别我们定义RPC⽅法的两个参数。
func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
