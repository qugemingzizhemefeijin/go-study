package main

import "fmt"

// Channel通信是在Goroutine之间进行同步的主要方法。在无缓存的Channel上的每一次发送操作都有其对应的接收操作想配对。
// 发送和接收操作通常发生在不同的Goroutine上（在同一个Goroutine上执行两个操作容易导致死锁）。
// 无缓存的Channel上的发送操作总在对应的接收操作完成前发生。（重点）

var done = make(chan bool)
var done2 = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好，世界"
	// done <- true
	close(done)
}

// 这个例子可保证打印出msg，该程序首先对msg进行写入，然后在done管道 上发送同步信号，随后从done接收对应的同步信号，最后执行print函数
// 若在关闭Channel后继续从中接收数据，接收者就会收到该Channel返回的零值。因此在这个例子中，用close关闭管道替代done<-true，依然能保证程序产生相同的行为。

// 对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。
// 基于上面的规则，交换两个Goroutine中的接收和发送操作也是可以的。（但是很危险）

func bGoroutine() {
	msg = "你好，世界2"
	<-done2
}

var limit = make(chan int, 3)

func main() {
	go aGoroutine()
	<-done
	fmt.Println(msg)

	// main线程中done<-true发送完成前，后台线程<-done接收已经开始，这保证了msg赋值被执行了，之后被打印了。
	// 后台线程首先对msg进行写入，然后从done中接收信号，随后main线程向done发送对应的信号，最后执行print函数。

	// 但是，若该Channel为带缓冲的，如 done=make(chan bool, 1)，main线程的done <- true接收操作将不会被后台线程的<-done接收操作阻塞，
	// 该程序将无法保证打印出"hello, world"
	go bGoroutine()
	done2 <- true
	fmt.Println(msg)
}
