package main

import (
	"fmt"
	"sync"
)

func main() {
	// 下面是错误版本，因为 mu.Lock() 和 mu.Unlock() 并不在同⼀个Goroutine中， 所以也就不满⾜顺序⼀致性内存模型。
	/*var mu sync.Mutex

	go func() {
		fmt.Println("Hello, World")
		mu.Lock()
	}()

	mu.Unlock()*/

	// 正确版本
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("Hello, World")
		mu.Unlock()
	}()

	mu.Lock()

	// 修复的⽅式是在 main 函数所在线程中执⾏两次 mu.Lock() ， 当第⼆次加锁时会因为锁已经被占⽤（不是递归锁） ⽽阻塞，
	// main 函数的阻塞状态驱动后台线程继续向前执⾏。 当后台线程执⾏到 mu.Unlock() 时解锁， 此时打印⼯作已经完成了，
	// 解锁会导致 main 函数中的第⼆个 mu.Lock() 阻塞状态取消， 此时后台线程和主线程再没有其它的同步事件参考，
	// 它们退出的事件将是并发的：在 main 函数退出导致程序退出时， 后台线程可能已经退出了， 也可能没有退出。
	// 虽然⽆法确定两个线程退出的时间， 但是打印⼯作是可以正确完成的

	// 也可以通过无缓存的管道来实现同步

	done := make(chan int)

	go func() {
		fmt.Println("你好，世界")
		<-done
	}()

	done <- 1

	// 上⾯的代码虽然可以正确同步， 但是对管道的缓存⼤⼩太敏感： 如果管道有缓存的话， 就⽆法保证main退出之前后台线程能正常打印了。
	// 更好的做法是将管道的发送和接收⽅向调换⼀下， 这样可以避免同步事件受管道缓存⼤⼩的影响

	done2 := make(chan int, 1)	// 带缓存的管道
	go func() {
		fmt.Println("你好，世界")
		done2 <- 1
	}()

	<- done2
}
