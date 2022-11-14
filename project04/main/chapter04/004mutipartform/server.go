package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// 这个可以获取表单格式为 multipart/form-data 的数据
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)

	// 此方法在需要的时候会自动调用ParseForm 或者 ParseMultipartForm，所以不需要前面主动的调用了
	// FormValue 方法只会获取第一个值，如果相同名称的表单话。
	// r.FormValue("hello")
	// r.PostFormValue("hello")

	// 如果表单类型为mulitpart/form-data，则只能使用r.MultipartForm来获取数据
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
