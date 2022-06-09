package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	lock sync.Mutex
)

//Queue ...
type Queue struct {
	count int   //保存上限
	array []int //数据
	size  int   //当前数据数量
}

func (q *Queue) push(val int) (err error) {
	lock.Lock()
	defer lock.Unlock()

	if q.size == q.count {
		err = errors.New("queue is full")
		return
	}
	q.array = append(q.array, val)
	q.size++
	return
}

func (q *Queue) pop() (val int, err error) {
	lock.Lock()
	defer lock.Unlock()

	if q.size == 0 {
		//err = errors.New("queue is empty")
		return
	}
	val = q.array[0]
	q.array = q.array[1:]
	q.size--
	return
}

//开启协程消费
func consume(i int, p *Queue) {
	for {
		time.Sleep(time.Duration((rand.Intn(5)+1)*100) * time.Millisecond) // 100-500毫秒随机
		val, err := p.pop()
		if err != nil {
			fmt.Println("err=", err)
		} else {
			if val > 0 {
				fmt.Printf("%d号协程服务 ---> %d号客户\n", i, val)
			}
		}
	}
}

//开启协程生产
func produce(p *Queue) {
	val := 1
	for {
		time.Sleep(time.Duration((rand.Intn(5)+1)*100) * time.Millisecond) // 100-500毫秒随机
		p.push(val)
		val++
	}
}

//1. 创建一个数组模拟队列，每隔一定时间[随机的]，给该数组添加一个数
//2. 启动两个协程，每隔一定时间(时间随机)到队列取出数据
//3. 在控制台输出
//   x号协程服务 ---> x号客户
//   x号协程服务 ---> x号客户
//   x号协程服务 ---> x号客户
//4. 使用锁机制即可
func main() {
	//初始化队列
	queue := &Queue{
		count: 10,
	}

	go produce(queue)

	for i := 0; i < 3; i++ {
		go consume(i+1, queue)
	}

	time.Sleep(24 * time.Hour)
}
