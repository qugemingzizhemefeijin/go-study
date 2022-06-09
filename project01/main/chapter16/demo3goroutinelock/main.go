package main

import (
	"fmt"
	_ "runtime"
	"sync"
	"time"
)

//1.编写一个函数，来计算各个数的阶乘，并放入到map中.
//2.我们启动多个协程，统计结果放入到map中.
//3.map应该做出一个全局的.

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁
	//sync 是包名：synchronized
	//Mutes: 互斥
	lock sync.Mutex
)

//test 计算n的阶乘，将这个结果放入到myMap中
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//这里我们将res放入到myMap中
	//加锁
	lock.Lock()
	//解锁
	defer lock.Unlock()
	myMap[n] = res
}

func main() {
	//需要计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来。要求使用goroutine来完成

	//分析思路：
	//1. 使用goroutine来完成，效率高，但是会出现并发/并行安全问题
	//2. 这里就他ichu了不同goroutine如何通信的问题

	//代码实现：
	//1. 使用goroutine来完成（看看使用goroutine并发完成会出现什么问题？然后我们会去解决）
	//2. 在运行某个程序时，如何知道是否存在资源竞争问题。方法很简单，在编译该程序时，增加一个参数 -race即可

	//加上这个就不会报错了 concurrent map writes [治标不治本]
	//runtime.GOMAXPROCS(1)

	// 我们这里开启多个协程完成这个任务
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)

	//go build -race main.go
	//./main.exe //可以看到竞争提醒

	lock.Lock()
	//这里理论上不需要加互斥锁，因为10秒可能都执行完毕了。但是加-race确实还是发现有这个竞争警告。
	//可能是主线程并不知道我们执行完毕，底层还是出现了资源争夺，加入锁后即可消除警告
	//这里我们输出结果，遍历这个结果
	for k, v := range myMap {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
	lock.Unlock()
	fmt.Println("OK")
}
