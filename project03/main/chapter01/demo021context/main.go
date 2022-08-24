package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 我们通过 close 来关闭 cancel 管道向多个Goroutine⼴播退出的指令。 不过这个程序依然不够稳健：
// 当每个Goroutine收到退出指令退出时⼀般会进⾏⼀定的清理⼯作， 但是退出的清理⼯作并不能保证被完成，
// 因为 main 线程并没有等待各个⼯作Goroutine退出⼯作完成的机制。 我们可以结合 sync.WaitGroup 来改进。

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		case <- ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("Hello")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	var wg sync.WaitGroup
	for i:=0;i<10;i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	time.Sleep(time.Second)
	cancel()

	wg.Wait()

	// 当并发体超时或 main 主动停⽌⼯作者Goroutine时， 每个⼯作者都可以安全退出
}
