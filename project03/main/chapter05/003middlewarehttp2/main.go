package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

// 使用中间件剥离非业务逻辑..

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello my honey!"))
}

// 这样就非常轻松地实现了业务与非业务之间的剥离，魔法就在于这个 timeMiddleware。
// 可以从代码中看到，我们的 timeMiddleware() 也是一个函数，其参数为 http.Handler，http.Handler 的定义在 net/http 包中。
// 任何方法实现了 ServeHTTP，即是一个合法的 http.Handler。
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		logger.Println(timeElapsed)
	})
}

func main() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	// https://github.com/gin-gonic/contrib
}