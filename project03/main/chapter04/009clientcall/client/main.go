package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net/rpc"
)

// Go 语言的 RPC 库最简单的使用方式是通过 Client.Call 方法进行同步阻塞调用。
// 执行异步调用的 Client.Go 方法实现如下：
/*
func (client *Client) Go(
	serviceMethod string, args interface{},
	reply interface{},
	done chan *Call,
) *Call {
	call := new(Call)
	call.ServiceMethod = serviceMethod
	call.Args = args
	call.Reply = reply
	call.Done = make(chan *Call, 10) // buffered.

	client.send(call)
	return call
}

首先是构造一个表示当前调用的 call 变量，然后通过 client.send 将 call 的完整参数发送到 RPC 框架。
client.send 方法调用是线程安全的，因此可以从多个 Goroutine 同时向同一个 RPC 链接发送调用指令。

当调用完成或者发生错误时，将调用 call.done 方法通知完成：
func (call *Call) done() {
	select {
	case call.Done <- call:
		// ok
	default:
		// We don't want to block here. It is the caller's responsibility to make
		// sure the channel has enough buffer space. See comment in Go().
	}
}

从 Call.done 方法的实现可以得知 call.Done 管道会将处理后的 call 返回。
 */
func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	// 前面的例子是将client封装到对象中，然后直接调用远程方法的。。这个地方的例子好比是dubbo的泛化调用的意思。其实就是001中的例子
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal("", err)
	}

	fmt.Println(reply)
}
