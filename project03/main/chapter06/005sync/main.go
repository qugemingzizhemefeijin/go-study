package main

import "sync"

// 要得到正确的结果的话，要把对计数器（counter）的操作代码部分加上锁

var counter int

func main() {
	var wg sync.WaitGroup
	var l sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}

	wg.Wait()
	println(counter)
}
