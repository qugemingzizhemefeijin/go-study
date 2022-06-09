package main

import (
	"context"
	"fmt"
)

// Go语⾔是带内存⾃动回收特性的， 因此内存⼀般不会泄漏。 在前⾯素数筛的例⼦中， GenerateNatural 和 PrimeFilter 函数内部都启动了新的Goroutine，
// 当 main 函数不再使⽤管道时后台Goroutine有泄漏的⻛险。 我们可以通过 context 包来避免这个问题， 下⾯是改进的素数筛实现：

func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()

	return ch
}

func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()

	return out
}

func main() {
	// 通过Context控制后台Goroutine状态
	ctx, cancel := context.WithCancel(context.Background())

	ch := GenerateNatural(ctx)
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ctx, ch, prime) // 基于新素数构造的过滤器
	}

	cancel()
}
