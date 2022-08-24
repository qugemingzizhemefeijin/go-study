package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"sync"
	"time"
)

// 在分布式场景下，我们也需要这种 “抢占” 的逻辑，这时候怎么办呢？我们可以使用 Redis 提供的 setnx 命令

func incr() {
	// context.WithTimeout 可用于一定时间关闭函数或调用，如调用超时，主动关闭调用
	// 在调用WithTimeout时，内部调用的是WithDeadline，实例化timeCtx，而它的父类是cancelCtx.
	// 使用time.AfterFunc，在一定时间后关闭通道，达到一定时间后通过管道通信，来关闭函数.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := client.SetNX(ctx, lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter++
	getResp := client.Get(ctx, counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(ctx, counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error1")
		}
	}
	println("current counter is ", cntValue)

	delResp := client.Del(ctx, lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}

	wg.Wait()
}
