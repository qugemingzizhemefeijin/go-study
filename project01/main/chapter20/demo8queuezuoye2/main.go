package main

import (
	"errors"
	"fmt"
)

//UINTMAX ...
const UINTMAX = ^uint(0)

//Node ...
type Node struct {
	val  int
	prev *Node
	next *Node
}

//Queue ...
type Queue struct {
	size uint
	head *Node
	tail *Node
}

func (q *Queue) push(val int) (err error) {
	if UINTMAX == q.size {
		err = errors.New("queue is full")
		return
	}

	last := q.tail //最后的节点

	//否则添加到队列中
	node := &Node{
		val:  val,
		prev: last,
	}

	q.tail = node    //将最后的节点设置为当前节点
	if last == nil { //如果最后的节点为空，则头节点也设置为当前节点
		q.head = node
	} else {
		last.next = node
	}
	q.size++
	return
}

func (q *Queue) pop() (val int, err error) {
	if q.size == 0 {
		err = errors.New("queue is empty")
		return
	}

	//从头部弹出一个
	head := q.head
	val = head.val
	next := head.next

	q.head = next

	if next == nil {
		q.tail = nil
	} else {
		next.prev = nil
	}
	q.size--
	return
}

func (q *Queue) show() {
	if q.size == 0 {
		fmt.Println("队列是空的~")
		return
	}
	for temp := q.head; temp != nil; temp = temp.next {
		fmt.Println(temp.val)
	}
	fmt.Println("===================")
}

//创建一个无界队列链表模拟队列，实现数据入队列，数据出队列，显示队列
func main() {
	queue := &Queue{}
	queue.push(1)
	queue.push(2)
	queue.push(3)
	queue.show()

	queue.pop()
	queue.show()
	queue.pop()
	queue.pop()
	queue.show()
}
