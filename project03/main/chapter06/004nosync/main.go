package main

import "sync"

// 在单机程序并发或并行修改全局变量时，需要对修改行为加锁以创造临界区

var counter int

func main() {
	var wg sync.WaitGroup
	for i := 0; i<1000;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	println(counter)
}

// 多次运行会得到不同的结果：
//
// 945
// 937
// 959
