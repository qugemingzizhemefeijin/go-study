package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// ServerMux的一个缺陷是无法使用变量实现URL模式匹配。
// 市面上有很多优秀的多路复用器可供使用。比如Gorilla Toolkit。或者本节介绍的HttpRouter

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

