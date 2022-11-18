package main

import (
	"html/template"
	"net/http"
	"os"
	"time"
)

// 获取工作目录
var WORK = getWordPath()

func getWordPath() (dir string) {
	dir, _ = os.Getwd()
	dir = dir + "/crs/project04/main/chapter05/006pipeline"
	return
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tmpl.html").Funcs(funcMap) // 定义一个模板
	// 在调用ParseFiles函数时，如果用户没有为模板文件中的模板定义名字，那么函数函数将使用模版文件的名字作为模板的名字
	t, _ = t.ParseFiles(WORK + "/tmpl.html")
	_ = t.Execute(w, time.Now())
}

// 测试界面管道
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}
