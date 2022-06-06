package main

import "fmt"

//CatNode ...
type CatNode struct {
	no   int
	name string
	next *CatNode
}

//InsertCatNode ...
func InsertCatNode(head *CatNode, newNode *CatNode) {
	//判断是不是第一次添加
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.next = head //形成环形
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	temp.next = newNode
	newNode.next = head
}

//ListCircleLink 输出这个环形的链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下...")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表...")
		return
	}

	for {
		fmt.Printf("猫的名字[id=%d,name=%s] -> \n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//DeleteCatNode ...
func DeleteCatNode(head *CatNode, no int) *CatNode {
	//1.先让temp指向head
	//2.让helper指向环形链表最后
	//3.让temp和要删除的id进行比较，如果相同，则helper完成删除[这里必须考虑如果删除就是头结点]

	temp := head
	helper := head //一旦找到要删除的节点，则使用helper来干掉节点

	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，无法删除")
		return head
	}

	//如果只有一个节点
	if temp.next == head { //说明当前环形链表只有一个节点
		if temp.no == no {
			temp.name = ""
			temp.no = 0
			temp.next = nil
		} else {
			fmt.Println("没有找到no=", no, "的节点")
		}
		return head
	}

	//将helper定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//如果有多个节点
	flag := true
	for {
		if temp.next == head { //如果到这里，说明我比较到最后一个[最后一个还没比较]
			break
		}
		if temp.no == no {
			if temp == head { //说明删除的是头结点
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("猫猫=%d 被删除了\n", no)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}

	if flag { //如果flag为真，则我们还需要再比较一次[因为在for循环上没有找到]
		if temp.no == no {
			helper.next = temp.next
			fmt.Printf("猫猫=%d 被删除了\n", no)
		} else {
			fmt.Println("对不起，没有no=", no)
		}
	}

	return head
}

//环形链表
func main() {
	//初始化一个环形链表的头节点
	head := &CatNode{}

	//创建一个节点
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "jerry",
	}
	cat3 := &CatNode{
		no:   3,
		name: "mark",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)
	fmt.Println("===删除")
	head = DeleteCatNode(head, 2)
	head = DeleteCatNode(head, 1)
	//head = DeleteCatNode(head, 3)
	ListCircleLink(head)
}
