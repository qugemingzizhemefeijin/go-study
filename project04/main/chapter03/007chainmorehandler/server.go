package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {

}

func (h HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO
		h.ServeHTTP(w, r)
	})
}

// 如果与HelloHandler处理器绑定的URL是/hello/而不是/hello，那么当浏览器请求/hello/there的时候，服务器找不到完全匹配的处理器时，
// 就会退而求其次，开始寻找能够与/hello/匹配的处理器，并最终找到helloHandler处理器。
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))

	server.ListenAndServe()
}
