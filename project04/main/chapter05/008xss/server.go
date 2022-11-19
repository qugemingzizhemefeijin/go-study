package main

import (
  "net/http"
  "html/template"
  "os"
)

// 获取工作目录
var WORK = getWordPath()

func getWordPath() (dir string) {
  dir, _ = os.Getwd()
  dir = dir + "/crs/project04/main/chapter05/008xss"
  return
}

// 上下文感知功能不仅能够自动对HTML转义，它还能防止基于JS，CSS甚至URL的XSS攻击。
func process(w http.ResponseWriter, r *http.Request) {
  // IE浏览器可以通过下面的头信息来关闭内置的XSS防御功能。
  w.Header().Set("X-XSS-Protection", "0")
  t, _ := template.ParseFiles(WORK + "/tmpl.html")
  _ = t.Execute(w, r.FormValue("comment"))
  // 如果想不对HTML进行转义，可以如下操作。
  _ = t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {  
  t, _ := template.ParseFiles(WORK + "/form.html")
  _ = t.Execute(w, nil)
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/process", process)
  http.HandleFunc("/", form)
  _ = server.ListenAndServe()
}
