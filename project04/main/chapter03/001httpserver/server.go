package main

import "net/http"

// 启动一个http服务
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}

	server.ListenAndServe()
}
