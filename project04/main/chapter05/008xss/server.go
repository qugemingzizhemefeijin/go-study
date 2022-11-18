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

func process(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("X-XSS-Protection", "0")
  t, _ := template.ParseFiles(WORK + "/tmpl.html")
  _ = t.Execute(w, r.FormValue("comment"))
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
