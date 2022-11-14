package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
	// 如果使用下面的方法获取文件的话，就不需要调用ParseMultipartForm
	file, _, err = r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintf(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
