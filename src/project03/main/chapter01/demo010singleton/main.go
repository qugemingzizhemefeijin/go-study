package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 原子操作配合互斥锁可以实现非常搞笑的单例模式。互斥锁的代价比普通整数的原子读写高很多，
// 在性能敏感的地方可以增加一个数字型的标志位，通过原子监测标志位状态降低互斥锁的使用次数来提高性能。

// singleton desc
type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

// Instance 单例
func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}

	return instance
}

// Once 还可以将代码抽取出来，做一个通过的单例模式
type Once struct {
	m    sync.Mutex
	done uint32
}

// Do 单例
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

var (
	s    *singleton
	once Once
)

// GetInstance DESC
func GetInstance() *singleton {
	once.Do(func() {
		s = &singleton{}
	})
	return s
}

// sync/atomic 包对基本的数值类型及复杂对象的读写都提供了原子操作的支持。atomic.Value原子对象提供了Load和Store两个原子方法
// 分别用于加载和保存数据，返回值和参数都是interface{}类型，因此可以用于任意的自定义复杂类型

func loadConfig() map[string]string {
	// 从数据库或者文件系统中读取配置信息，然后以map的形式存放在内存里
	return make(map[string]string)
}

func requests() chan int {
	// 将从外界中接受到的请求放入到channel里
	return make(chan int)
}

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	var config atomic.Value // 保存你当前配置信息

	// 初始化配置信息
	config.Store(loadConfig())

	// 启动一个后台线程，加载更新后的配置信息
	go func() {
		for {
			time.Sleep(10 * time.Second)
			config.Store(loadConfig())
		}
	}()

	// 用于处理请求的工作者线程始终采用最新的配置信息
	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				// 对应于取值操作 c := config
				// 由于Load()返回的是一个interface{}类型，所以我们要先强制转换一下
				c := config.Load().(map[string]string)
				// 这里是根据配置信息处理请求的逻辑...
				_, _ = r, c
			}
		}()
	}

	fmt.Println("==================")
	// go 语言也有可见性这一说

	// 下面这一段代码可能会让程序陷入死循环
	/*go setup()
	for !done {}
	print(a)*/

	// 在Go语言中，同一个Goroutine线程内部，顺序一致性内存模型是得到保证的。但是不同的Goroutine之间，并不满足顺序一致性内存模型。
	// 需要通过明确定义的同步事件来作为同步的参考。如果两个事件不可排序，那么就说这两个事件是并发的。为了最大化并行，Go语言的编译器和处理器
	// 在不影响上述规则的前提下可能会对执行语句重新排序（CPU也会对一些指令进行乱序执行）

	// go println("你好，世界")
	// 根据规范，main函数退出时程序结束，不会等待任何后台线程。因为Goroutine的执行和main函数的返回事件是并发的，谁都有可能先发生。

	// 解决办法如下：
	done := make(chan int)

	go func() {
		fmt.Println("你好，世界")
		done <- 1
	}()

	<-done

	// 当<-done执行时，必然要求done<- 1也已经执行。根据同一个Goroutine依然满足顺序一致性规则，我们可以判断当done<- 1执行时，print语句必然执行完毕

	// 也可以通过互斥量来完成
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("你好，世界")
		mu.Unlock()
	}()

	mu.Lock()

	// 可以确定后台线程的mu.Unlock()必然在print完成之后发生，main函数的第二个lock必然在后台线程的unlock之后发生，此时后台线程的打印工作已经顺利完成。
}
