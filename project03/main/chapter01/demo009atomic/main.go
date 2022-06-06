package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total struct {
	sync.Mutex
	value int
}

// 使用互斥锁来保护共享资源
func worker(wg *sync.WaitGroup) {
	defer wg.Done() // 解锁

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

// 还可以使用atomic，性能比锁要好
var tt uint64

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&tt, i)
	}
}

// 原子操作
func main() {
	var wg sync.WaitGroup
	wg.Add(2) // 此功能有点像java的CountDownLatch
	go worker(&wg)
	go worker(&wg)
	wg.Wait() // 等待Done完成

	fmt.Println(total.value)

	fmt.Println("===============")
	var aa sync.WaitGroup
	aa.Add(2)

	go worker2(&aa)
	go worker2(&aa)
	aa.Wait()
	fmt.Println(tt)
}
