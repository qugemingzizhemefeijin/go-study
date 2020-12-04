package main

import (
	"fmt"
	"math"
	"time"
)

//要求统计1-8000的数字中，哪些是素数？
//传统的方法，就是使用一个循环，循环的判断各个数是不是素数
//使用并发/并行的方式，将统计素数的任务分配给多个(4个)goroutine去完成，完成任务时间短。

//向 intChan放入 1 - 8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}

	//关闭channel
	close(intChan)
}

//从intChan取出数据，并判断是否为素数，如果是，就放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		//判断num是否是素数
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum协程退出了")
	exitChan <- true
}

func main() {
	num := 4
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000) //放入结果
	//标识退出的管道
	exitChan := make(chan bool, num)

	start := time.Now().Unix()

	go putNum(intChan)
	//开启4个协程
	for i := 0; i < num; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < num; i++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)
		//当我们从exitChan取出了4个结果，就可以放心的关闭primeNum
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(res)
	}

	fmt.Println("exit success")
}
