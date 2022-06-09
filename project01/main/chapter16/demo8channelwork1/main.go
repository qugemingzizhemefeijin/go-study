package main

import (
	"fmt"
	"sync"
	"time"
)

//1.启动一个协程，将1-2000的数放入到一个channel中，比如numChan
//2.启动8个协程，从numChan取出数（比如n），并计算1+...+n的值，并存放到resChan
//3.最后8个协程协同完成工作后，再遍历resChan，显示结果[如 res[1]=1...res[10]=55..]
//4.注意：考虑resChan chan int 是否合适？

func produceNum(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func calcNum(numChan chan int, resChan *[2000]int, wait *sync.WaitGroup) {
	for {
		n, ok := <-numChan
		if !ok {
			break
		}

		//fmt.Println(n)

		total := 0
		for i := 1; i <= n; i++ {
			total += i
		}
		(*resChan)[n-1] = total

		//模拟计算很慢
		time.Sleep(time.Millisecond * 5)
	}
	wait.Done()
}

func main() {
	n, m := 2000, 8
	var wait sync.WaitGroup //声明线程等待对象
	numChan := make(chan int, n)
	var resChan [2000]int

	//启动生成协程
	go produceNum(numChan)
	//启动8个计算协程
	for i := 0; i < m; i++ {
		go calcNum(numChan, &resChan, &wait)
	}

	wait.Add(m)
	wait.Wait() //主线程等待其他协程完成任务后，再结束

	//遍历resChan
	for i := 0; i < n; i++ {
		fmt.Printf("res[%d]=%d\n", i+1, resChan[i])
	}

	fmt.Println("exit success")
}
