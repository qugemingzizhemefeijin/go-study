package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {

}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// 模拟命令
// curl -s localhost:1234/jsonrpc -X POST --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
// 返回： {"id":0,"result":"hello:hello","error":null}
func main() {
	_ = rpc.RegisterName("HelloService", new(HelloService))

	// RPC的服务架设在“/jsonrpc”路径
	// 在处理函数中基于http.ResponseWriter和http.Request类型的参数构造⼀个io.ReadWriteCloser类型的conn通道。
	// 然后基于conn构建针对服务端的json编码解码器
	// 最后通过rpc.ServeRequest函数为每次请求处理⼀次RPC⽅法调⽤
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer: w,
		}

		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	_ = http.ListenAndServe(":1234", nil)
}
