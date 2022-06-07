package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 生产者，消费者模型

// 生产者：生成factor整数倍的序列
func Producer(factor int, out chan<- int) {
	for i:=0;;i++ {
		out <- i*factor
	}
}

// 消费者
func Consumer(in <- chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)

	go Producer(3, ch) // 生成3个倍数的序列
	go Producer(5, ch) // 生成5倍数的序列
	go Consumer(ch) // 消费生成的队列

	// 运行一段时间后退出
	// time.Sleep(5 * time.Second)

	// 我们可以让 main 函数保存阻塞状态不退出， 只有当⽤户输⼊ Ctrl-C 时才真正退出程序
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
