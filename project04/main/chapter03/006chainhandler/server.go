package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// log函数接受一个HandlerFunc类型的函数作为参数，然后返回另一个HandlerFunc类型的函数作为值。
// 因为hello函数就是一个HandlerFunc类型的函数，所以代码log(hello)实际就是将hello函数发送至log函数以内。
// 换句话说，这段代码串联起了log函数和hello函数

// log函数的返回值是一个匿名函数，因为这个匿名函数接受一个ResponseWriter和一个Request指针作为参数，所以它实际也是一个HandlerFunc。
// 在匿名函数内部，程序首先会获取被传入的HandlerFunc的名字，然后再调用这个HandlerFunc。
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main () {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", log(hello))

	server.ListenAndServe()
}
