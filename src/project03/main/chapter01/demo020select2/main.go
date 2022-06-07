package main

import (
	"fmt"
	"sync"
	"time"
)

/*func worker(channel chan bool) {
	for {
		select {
		default:
			fmt.Println("Hello")
			// 正常工作
		case <- channel:
			// 退出
		}
	}
}*/

func worker(wg *sync.WaitGroup, channel chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <- channel:
			return
		}
	}
}

func main() {
	// 当有多个管道均可操作时， select 会随机选择⼀个管道。 基于该特性我们可以⽤ select 实现⼀个⽣成随机数序列的程序
	channel := make(chan bool)
	// go worker(channel)

	var wg sync.WaitGroup
	for i:=0;i<10;i++ {
		wg.Add(1)
		go worker(&wg, channel)
	}

	time.Sleep(time.Second)
	// channel <- true

	// 但是管道的发送操作和接收操作是⼀⼀对应的， 如果要停⽌多个Goroutine那么可能需要创建同样数量的管道， 这个代价太⼤了。
	// 其实我们可以通过 close 关闭⼀个管道来实现⼴播的效果， 所有从关闭管道接收的操作均会收到⼀个零值和⼀个可选的失败标志
	close(channel)

	// 我们通过 close 来关闭 cancel 管道向多个Goroutine⼴播退出的指令。 不过这个程序依然不够稳健：
	// 当每个Goroutine收到退出指令退出时⼀般会进⾏⼀定的清理⼯作， 但是退出的清理⼯作并不能保证被完成，
	// 因为 main 线程并没有等待各个⼯作Goroutine退出⼯作完成的机制。 我们可以结合 sync.WaitGroup 来改进
	wg.Wait()
}
