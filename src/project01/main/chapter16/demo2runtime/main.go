package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test = ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置多个CPU
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")

	//1 go1.8后，默认让程序运行在多个核上，可以不用设置
	//2 go1.8前，需要设置一下GOMAXPROCS，可以更高效的利用CPU

	//测试一下协程
	//runtime.GOMAXPROCS(2)
	//go test()
	//for i := 0; i < 10000000000000; i++ {
	//fmt.Println("main = ", i)
	//time.Sleep(time.Second)
	//}
}
