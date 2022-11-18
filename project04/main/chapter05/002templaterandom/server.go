package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// 获取工作目录
var WORK = getWordPath()

func getWordPath() (dir string) {
	dir, _ = os.Getwd()
	dir = dir + "/crs/project04/main/chapter05/002templaterandom"
	return
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(WORK + "/tmpl.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
