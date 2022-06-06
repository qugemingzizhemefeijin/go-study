package main

import (
	"fmt"
	"time"
)

var a string

func f() {
	fmt.Println(a)
}

func hello() {
	a = "hello, world"
	go f()
}

// 执行go f()语句创建goroutine和hello函数是在同一个Goroutine中执行，根据语句的书写顺序可以确定Goroutine的创建发生在hello函数返回之前。
// 但是新创建Goroutine对应的f()的执行事件和hello函数返回的事件则是不可排序的，也就是并发的。
// 调用hello可能会在将来的某一时刻打印"hello, world"，也可能是在hello函数执行完成后才打印
func main() {
	hello()
	time.Sleep(time.Second * 2)
}
