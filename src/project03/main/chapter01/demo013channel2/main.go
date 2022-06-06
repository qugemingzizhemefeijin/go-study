package main

import (
	"fmt"
	"time"
)

// 我们可以根据控制 Channel 的缓存大小来控制并发执行的 Goroutine 的最大数目

var limit = make(chan int, 3)
var work = []func(){
	func() { fmt.Println("1"); time.Sleep(1 * time.Second) },
	func() { fmt.Println("2"); time.Sleep(1 * time.Second) },
	func() { fmt.Println("3"); time.Sleep(1 * time.Second) },
	func() { fmt.Println("4"); time.Sleep(1 * time.Second) },
	func() { fmt.Println("5"); time.Sleep(1 * time.Second) },
}

// 在循环创建 Goroutine 过程中，使用了匿名函数并在函数中引用了循环变量 w，
// 由于 w 是引用传递的而非值传递，因此无法保证 Goroutine 在运行时调用的 w 与循环创建时的 w 是同一个值，
// 为了解决这个问题，我们可以利用函数传参的值复制来为每个 Goroutine 单独复制一份 w。

// 循环创建结束后，在 main 函数中最后一句 select{} 是一个空的管道选择语句，该语句会导致 main 线程阻塞，
// 从而避免程序过早退出。还有 for{}、<-make(chan int) 等诸多方法可以达到类似的效果。因为 main 线程被阻塞了，
// 如果需要程序正常退出的话可以通过调用 os.Exit(0) 实现。
func main() {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select {}
}
