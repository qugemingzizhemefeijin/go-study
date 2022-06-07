package main

import (
	"fmt"
	"time"
)

//Go协程的特点
//1.有独立的栈空间
//2.共享程序堆空间
//3.调度由用户控制
//4.协程是轻量级的线程

//编写一个函数，每隔1秒输出"hello,world"
func test() {
	for i := 1; i <= 20; i++ {
		fmt.Printf("hello,world, %d\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	//1.在主线程(可以理解成进程)中，开启一个goroutine，该协程每隔1秒输出"Hello,World!"
	//2.在主线程中也每隔1秒输出"Hello,golang!"，输出10次后，退出程序
	//3.要求主线程和goroutine同时执行
	go test() //开启了一个协程(这里会有问题哦，如果主线程执行完毕，虽然协程没有执行完，但是协程也被关闭了)

	//1.如果主线程退出了，则协程即使还没有执行完毕，也会退出
	//2.当然协程也可以在主线程没有退出前，就自己结束了，比如完成了自己的任务（主线程还是会继续完成未完成的事业）

	for i := 1; i <= 10; i++ {
		fmt.Printf("hello,golang, %d\n", i)
		time.Sleep(time.Second)
	}

	//得出结论：
	//1) 主线程是一个物理线程，直接作用在cpu上的。是重量级的，非常耗费CPU资源；
	//2) 协程从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗相对小；
	//3) Go的协程机制是重要的特点，可以轻松的开启上万个协程。
	//   其它编程语言的并发机制是一般基于线程的，开启过多的线程，资源耗费大，这里就突显Go在并发上的优势了

	//goroutine的调度模型（MPG）
	// 1. M: 操作系统的主线程(物理线程)
	// 2. P: 协程执行需要的上下文
	// 3. G: 协程

	//MPG模式运行的状态1
	//1) 当前程序有三个M，如果三个M都在一个CPU运行，就是并发，如果在不同的CPU运行就是并行
	//2) M1,M2,M3正在执行一个G，M1的协程队列有三个，M2的协程队列有三个，M3写成队列有两个
	//3) GO的协程是轻量级的线程，是逻辑态的，GO可以容易的起上万个协程
	//4) 其他程序C/Java的多线程，往往是内核态，比较重量级，几千个线程可能耗光CPU

	//MPG模式运行的状态2
	//1) 原来的情况是M0主线程正在执行Go协程，另外有三个协程在队列等待
	//2) 如果Go协程阻塞，比如读取文件或者数据库等
	//3) 这时就会创建M1主线程（也可能是从已有的线程池中取出M1），并且将等待的3个协程挂到M1下开始执行，M0的主线程下的Go任然执行文件IO的读写
	//4) 这样的MPG调度模式，可以既让GO执行，同时也不会让队列的其它协程一直阻塞，仍然而已并发/并行执行
	//5) 等到GO不阻塞了，M0会被放到空闲的主线程继续执行（从已有的线程池中取），同时GO又会被唤醒
}