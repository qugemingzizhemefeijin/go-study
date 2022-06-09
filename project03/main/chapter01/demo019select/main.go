package main

import "fmt"

// 当 select 有多个分⽀时， 会随机选择⼀个可⽤的管道分⽀， 如果没有可⽤的管道分⽀则选择 default 分⽀， 否则会⼀直保存阻塞状态
// 基于 select 实现的管道的超时判断：
/*
select {
	case v := <-in:
		fmt.Println(v)
	case <-time.After(time.Second):
		return // 超时
}

通过 select 的 default 分⽀实现⾮阻塞的管道发送或接收操作：
select {
	case v := <-in:
		fmt.Println(v)
	default:
		// 没有数据
}

通过 select 来阻⽌ main 函数退出：
func main() {
	// do some thins
	select{}
}
 */
func main() {
	// 当有多个管道均可操作时， select 会随机选择⼀个管道。 基于该特性我们可以⽤ select 实现⼀个⽣成随机数序列的程序
	ch := make(chan int)
	go func() {
		for {
			select{
				case ch <- 0:
				case ch <- 1:
			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
