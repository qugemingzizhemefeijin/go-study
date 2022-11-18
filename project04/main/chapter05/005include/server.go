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
	dir = dir + "/crs/project04/main/chapter05/005include"
	return
}

// Hello World会被传递给t1.html
// 但是因为t1.html没有给t2.html传递参数，所以t2.html中的 {{ . }} 就为空
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(WORK + "/t1.html", WORK + "/t2.html")
	_ = t.Execute(w, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}
