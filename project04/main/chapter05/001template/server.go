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
	dir = dir + "/crs/project04/main/chapter05/001template"
	return
}

func process(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(WORK + "/tmpl.html"))
	t.Execute(writer, "Hello World!")
}

func main() {
	//  tmpl, err := template.New("").Delims("[[", "]]").ParseFiles("base.tmpl", "homepage/inner.tmpl")
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
