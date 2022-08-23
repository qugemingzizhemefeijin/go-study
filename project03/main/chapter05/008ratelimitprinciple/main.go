package main

import (
	"fmt"
	"time"
)

// 令牌桶模型实际上就是对全局计数的加减法操作过程， 但使⽤计数需要我们⾃⼰加读
// 写锁， 有⼩⼩的思想负担。 如果我们对Go语⾔已经⽐较熟悉的话， 很容易想到可以⽤buffered channel
// 来完成简单的加令牌取令牌操作

var fillInterval = time.Millisecond * 10
var capacity = 10
var tokenBucket = make(chan struct{}, capacity)

// 自己造的限流的轮子就是用这种方式来实现的，不过如果开源版 ratelimit 也如此的话，那我们也没什么可说的了。现实并不是这样的。
// 其实主要是性能比较差，可以使用 github.com/juju/ratelimit 来惰性求值。
func TakeAvailable(block bool) bool {
	var takenResult bool
	if block {
		select {
		case <- tokenBucket:
			takenResult = true
		}
	} else{
		select {
		case <- tokenBucket:
			takenResult = true
		default:
			takenResult = false
		}
	}
	return takenResult
}

func main() {
	// 在 1s 钟的时候刚好填满 100 个，没有太大的偏差。不过这里可以看到，Go 的定时器存在大约 0.001s 的误差，
	// 所以如果令牌桶大小在 1000 以上的填充可能会有一定的误差。

	fillToken := func() {
		ticker := time.NewTicker(fillInterval)
		for {
			select {
			case <- ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
				}
				fmt.Println("current token cnt:", len(tokenBucket), time.Now())
			}
		}
	}

	go fillToken()

	time.Sleep(time.Hour)
}
