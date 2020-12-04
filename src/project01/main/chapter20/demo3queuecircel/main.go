package main

import (
	"errors"
	"fmt"
)

//CircleQueue 使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int //5
	array   [5]int
	head    int //队首(含队首元素)
	tail    int //队尾(不含队尾元素)
}

//Push 入队
func (q *CircleQueue) Push(val int) (err error) {
	if q.IsFull() {
		err = errors.New("queue full")
		return
	}

	q.array[q.tail] = val
	q.tail++
	if q.tail == q.maxSize {
		q.tail = 0
	}
	return
}

//Pop 出队
func (q *CircleQueue) Pop() (val int, err error) {
	if q.IsEmpty() {
		err = errors.New("queue empty")
		return
	}

	val = q.array[q.head]
	q.head++
	if q.head == q.maxSize {
		q.head = 0
	}
	return
}

//IsEmpty 判断是否为空
func (q *CircleQueue) IsEmpty() bool {
	return q.head == q.tail
}

//IsFull 判断是否满了
func (q *CircleQueue) IsFull() bool {
	return (q.tail+1)%q.maxSize == q.head
}

//Size 获取队列元素数量
func (q *CircleQueue) Size() int {
	return (q.tail + q.maxSize - q.head) % q.maxSize
}

//Show 显示队尾
func (q *CircleQueue) Show() {
	size := q.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}
	i := q.head
	for i != q.tail {
		fmt.Printf("array[%d]=%d\t", i, q.array[i])
		// if i == size-1 {
		// 	i = 0
		// } else {
		// 	i++
		// }
		i = (i + 1) % q.maxSize
	}
	fmt.Println()
}

func main() {
	queue := &CircleQueue{
		maxSize: 5,
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
			err := queue.Push(val)
			if err != nil {
				fmt.Println("err=", err)
			} else {
				fmt.Println("加入队列OK")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println("err=", err)
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.Show()
		case "exit":
			loop = false
		default:
			fmt.Println("你输入的有误，请重新输入")
		}
	}

	fmt.Println("exit success")
}
