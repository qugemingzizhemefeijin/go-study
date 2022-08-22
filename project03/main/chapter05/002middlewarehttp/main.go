package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

func hello(wr http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	wr.Write([]byte("hello golang!"))
	timeElapsed := time.Since(timeStart)
	logger.Println(timeElapsed)
}

func showFriendsHandler(wr http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	wr.Write([]byte("your friends is tom and jerry!"))
	timeElapsed := time.Since(timeStart)
	logger.Println(timeElapsed)
}

// 这个例子可以看出如果将所有的请求都统计耗时，代码耦合度太高，要是需要修改统计信息的话需要全部都要改一遍。
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/friends/show", showFriendsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}