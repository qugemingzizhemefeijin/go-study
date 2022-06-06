package main

import (
	"fmt"
	"sync"
)

// 基于带缓存的管道，实现10个后台线程分别打印：
func main() {
	done := make(chan int, 10) // 带10个缓存

	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("你好，世界")
			done <- 1
		}()
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}

	fmt.Println("==============")

	// 对于这种要等待 N 个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用 sync.WaitGroup 来等待一组事件：
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好，世界")
			wg.Done()
		}()
	}

	// 等待 N 个后台线程完成
	wg.Wait()
}
