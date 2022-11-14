package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User 	string
	Threads []string
}

// 接收Request结构指针的原因是为了让服务器能够察觉到处理器对Request结构的修改
// ResponseWriter实际上是response这个非导出结构的接口，而ResponseWriter在使用response结构时，传递的也是指向response结构的指针
// 也就是说ResponseWriter是以传引用而不是传值的方式在使用response结构。

// 换句话说，实际上ServeHTTP函数的两个参数传递的都是引用而不是值，虽然ResponseWriter看上去像是一个值，但它既是确实一个带有结构指针的接口

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Hello Golang!</title></head>
	<body><h1>Hello Golang!</h1></body>
	</html>`

	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such server, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post {
		User:	"jack chan",
		Threads: []string{"first", "second", "third"},
	}

	j, _ := json.Marshal(post)
	w.Write(j)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}
