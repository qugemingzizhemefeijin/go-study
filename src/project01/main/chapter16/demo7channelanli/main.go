package main

import (
	"fmt"
	"time"
)

//请完成goroutine和channel协同工作的案例，具体要求：
//1. 开启一个writeData协程，向管道intChan中写入50个整数
//2. 开启一个readData协程，从管道intChan中读取writeData写入的数据
//3. 注意：writeData和readData操作的是同一个管道
//4. 主线程需要等待writeData和readData协程都完成工作才能退出

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		//这里模拟慢慢写入
		fmt.Printf("writeData 写入到数据=%d \n", i)
		time.Sleep(time.Second)
	}
	close(intChan) //关闭
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok { //获取不到数据了，则ok=false
			break
		}
		fmt.Printf("readData 读取到数据=%d \n", v)
	}
	//任务完成，写入到主线程退出通道
	exitChan <- true
	close(exitChan) //关闭
}

func main() {
	//首先创建两个管理
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1) //用于让让主线程阻塞

	go writeData(intChan)
	go readData(intChan, exitChan)

	//让主线程等待任务完成后，背通知退出
	for {
		v, _ := <-exitChan
		if v {
			break
		}
	}

	fmt.Println("exit success")
}
