package main

import "sync"

// 在某些场景，我们只是希望一个任务有单一的执行者。而不像计数器场景一样，所有 goroutine 都执行成功。
// 后来的 goroutine 在抢锁失败后，需要放弃其流程。这时候就需要 trylock 了。

// trylock 顾名思义，尝试加锁，加锁成功执行后续流程，如果加锁失败的话也不会阻塞，而会直接返回加锁的结果。
// 在 Go 语言中我们可以用大小为 1 的 Channel 来模拟 trylock

// Lock try lock
type Lock struct {
	c chan struct{}
}

// NewLock generate a try lock
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	// 这里必须先向l.c中放入一个对象，否则所有的协程都无法拿到锁
	l.c <- struct{}{}
	return l
}

// Lock try lock, return lock result
func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:

	}
	return lockResult
}

// Unlock, Unlock the try lock
func (l Lock) Unlock() {
	l.c <- struct{}{}
}

var counter int

// 因为我们的逻辑限定每个 goroutine 只有成功执行了 Lock 才会继续执行后续逻辑，
// 因此在 Unlock 时可以保证 Lock 结构体中的 channel 一定是空，从而不会阻塞，也不会失败。
// 代码使用了大小为 1 的 channel 来模拟 trylock，理论上还可以使用标准库中的 CAS 来实现相同的功能且成本更低，读者可以自行尝试。

// 在单机系统中，trylock 并不是一个好选择。因为大量的 goroutine 抢锁可能会导致 CPU 无意义的资源浪费。有一个专有名词用来描述这种抢锁的场景：活锁。

// 活锁指的是程序看起来在正常执行，但 CPU 周期被浪费在抢锁，而非执行任务上，从而程序整体的执行效率低下。
// 活锁的问题定位起来要麻烦很多。所以在单机场景下，不建议使用这种锁。

func main() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.Unlock()
		}()
	}
	wg.Wait()
}
