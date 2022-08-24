package main

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

// 基于 ZooKeeper 的锁与基于 Redis 的锁的不同之处在于 Lock 成功之前会一直阻塞，这与我们单机场景中的 mutex.Lock 很相似。
func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		panic(err)
	}

	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err = l.Lock()
	if err != nil {
		panic(err)
	}
	println("lock succ, do your business logic")

	time.Sleep(time.Second * 10)

	fmt.Println("===================")

	// unlock
	l.Unlock()
	println("unlock success, finish business logic")
}
