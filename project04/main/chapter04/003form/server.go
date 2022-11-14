package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// 解析普通的表单提交
	r.ParseForm()
	// 获取所有的表单提交数据，包括GET
	// fmt.Fprintln(w, r.Form)
	// 获取指定表单的提交至
	//fmt.Fprintln(w, r.Form["hello"])

	// 只获取post
	fmt.Fprintln(w, r.PostForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
