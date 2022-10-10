package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// 创建一个默认的多路复用器
	mux := http.NewServeMux()
	// 指定目录中的静态文件服务的处理器
	files := http.FileServer(http.Dir(config.Static))
	// 这个当请求/static/，则自动去掉 /static/ 并去 /public 中查找目录
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// 启动服务，从配置文件中读取绑定的地址和端口
	server := &http.Server{
		Addr: config.Address,
		Handler: mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
