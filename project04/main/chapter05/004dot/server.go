package main

import (
	"html/template"
	"net/http"
	"os"
)

// 获取工作目录
var WORK = getWordPath()

func getWordPath() (dir string) {
	dir, _ = os.Getwd()
	dir = dir + "/crs/project04/main/chapter05/004dot"
	return
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(WORK + "/tmpl.html")
	_ = t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}
