package main

import (
	"errors"
	"fmt"
)

//Queue 使用一个结构体管理队列[此队列是与问题的，加满后就不能再使用了]
type Queue struct {
	maxSize int
	array   [5]int //数组=>模拟队列
	front   int    //表示指向队列列首 (不含队首)
	rear    int    //表示指向队列尾部(包含最后元素)
}

//AddQueue 添加数据到队列
func (q *Queue) AddQueue(val int) error {
	//先判断队列是否已满
	if q.rear == q.maxSize-1 {
		return errors.New("queue full")
	}

	q.rear++ //rear 后移
	q.array[q.rear] = val
	return nil
}

//ShowQueue 显示队列，找到队首，遍历到队尾
func (q *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	// front 不包含队首元素
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.array[i])
	}
	fmt.Println()
}

//GetQueue 从队列中取出数据
func (q *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if q.front == q.rear {
		err = errors.New("queue empty")
		return
	}
	q.front++
	val = q.array[q.front]

	return
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	var loop = true
	for loop {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列中获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示退出队列")

		fmt.Scanln(&key)

		switch key {
		case "add":
			fmt.Println("输入你要入队列的数：")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println("err=", err)
			} else {
				fmt.Println("加入队列OK")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println("err=", err)
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			loop = false
		default:
			fmt.Println("你输入的有误，请重新输入")
		}
	}

	fmt.Println("exit success")
}
