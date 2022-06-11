package main

import (
	"fmt"
	"log"
	"net/rpc"

	"../pb"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dial error:", err)
	}

	var args = pb.String{
		Value: "Hello, World",
	}
	var reply pb.String
	err = client.Call("HelloService.Hello", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())
}
