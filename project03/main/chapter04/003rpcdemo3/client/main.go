package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 请求的json数据对象在内部对应两个结构体： 客户端是clientRequest， 服务端是serverRequest。
// clientRequest和serverRequest结构体的内容基本是⼀致的：
/*
type clientRequest struct {
	Method string `json:"method"`
	Params [1]interface{} `json:"params"`
	Id uint64 `json:"id"`
}
type serverRequest struct {
	Method string `json:"method"`
	Params *json.RawMessage `json:"params"`
	Id *json.RawMessage `json:"id"`
}

在获取到RPC调⽤对应的json数据后， 我们可以通过直接向架设了RPC服务的TCP服务器发送json数据模拟RPC⽅法调⽤：
echo -e '{"method":"HelloService.Hello","params":["hello"],"id":1}' | nc localhost 1234
返回的结果也是⼀个json格式的数据：
{"id":1,"result":"hello:hello","error":null}

返回的json数据也是对应内部的两个结构体： 客户端是clientResponse， 服务端是serverResponse。两个结构体的内容同样也是类似的：
type clientResponse struct {
	Id uint64 `json:"id"`
	Result *json.RawMessage `json:"result"`
	Error interface{} `json:"error"`
}

type serverResponse struct {
	Id *json.RawMessage `json:"id"`
	Result interface{} `json:"result"`
	Error interface{} `json:"error"`
}
 */
func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dail:", err)
	}

	// 在确保客户端可以正常调⽤RPC服务的⽅法之后， 我们⽤⼀个普通的TCP服务代替Go语⾔版本的RPC服务， 这样可以查看客户端调⽤时发送的数据格式。
	// ⽐如通过nc命令 nc -l 1234 在同样的端⼝启动⼀个TCP服务。 然后再次执⾏⼀次RPC调⽤将会发现nc输出了以下的信息：
	// {"method":"HelloService.Hello","params":["hello"],"id":0}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
