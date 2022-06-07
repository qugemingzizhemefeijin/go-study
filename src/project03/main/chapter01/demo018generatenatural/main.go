package main

import "fmt"

// 并发版本素数筛的实现

// 返回生成自然数序列的管道：2, 3, 4
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器，删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
// 这里其实会启动100个Goroutine，第一个是2开始的数列生成器
// 第二个就是过滤%2=0的数列
// 第三个就是过滤%3=0的数列，以此类推
func main() {
	ch := GenerateNatural() // 自然数序列：2，3，4
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}
