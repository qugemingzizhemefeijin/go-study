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
	dir = dir + "/crs/project04/main/chapter05/007context_aware"
	return
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(WORK + "/tmpl.html")
	content := `I asked: <i>"What's up?"</i>`
	_ = t.Execute(w, content)
}

// 上下文感知，就是用于被现实的内容实施正确的转义：
// 如果模板显示的是html内容，那么模板将对其实施HTML转义。
// 如果模板显示的是JavaScript格式内容，那么将对其实施JavaScript转义；
// GO模板引擎还可以识别出内容中的URL或者CSS样式。
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	_ = server.ListenAndServe()
}
