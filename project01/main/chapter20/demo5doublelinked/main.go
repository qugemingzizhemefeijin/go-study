package main

import (
	"fmt"
)

//HeroNode 英雄排行榜
type HeroNode struct {
	no       int
	name     string
	nickname string
	prev     *HeroNode //指向上一个节点
	next     *HeroNode //这个表示指向下一个节点
}

//InsertHeroNode ...
func InsertHeroNode(head *HeroNode, node *HeroNode) {
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
	node.prev = temp
}

//InsertHeroNode2 给链表插入一个节点
//编写第二种方法，根据no的编号从小到大插入
func InsertHeroNode2(head *HeroNode, node *HeroNode) {
	//思路
	//1.要找到适当的节点

	temp := head
	flag := true

	for {
		if temp.next == nil { //表示找到最后了
			break
		} else if temp.next.no > node.no {
			//说明node应该插入到temp后面
			break
		} else if temp.next.no == node.no {
			//说明链表中已经有这个no，就不让插入
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("无法加入相同No的英雄")
		return
	}

	//2.找到最后一个节点，将最后一个节点的next指向新节点
	node.next = temp.next
	node.prev = temp

	if temp.next != nil {
		temp.next.prev = node
	}
	temp.next = node
}

//ListHeroNode 显示链表的所有节点信息
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}
	for temp = temp.next; temp != nil; temp = temp.next {
		fmt.Printf("[%d, %s, %s] ==> ", temp.no, temp.nickname, temp.name)
	}
	fmt.Println()
}

//ListHeroNode2 显示链表的所有节点信息
func ListHeroNode2(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}

	//让temp指向双向链表的最后
	for temp.next != nil {
		temp = temp.next
	}

	for ; temp != head; temp = temp.prev {
		fmt.Printf("[%d, %s, %s] <== ", temp.no, temp.nickname, temp.name)
	}
	fmt.Println()
}

//DeleteHeroNode 删除一个节点
func DeleteHeroNode(head *HeroNode, no int) {
	flag := false
	temp := head
	for ; temp != nil; temp = temp.next {
		if temp.next == nil {
			break
		} else if temp.next.no == no {
			flag = true
			break
		}
	}

	if flag {
		//直接删除即可，其实就是当前指针指向下一个指针的next即可
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.prev = temp
		}
	} else {
		fmt.Println("sorry,要删除的No不存在")
	}
}

//双向链表
func main() {
	//1.先创建一个头节点
	head := &HeroNode{}

	//2.创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "呼保义",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "吴用",
		nickname: "智多星",
	}

	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero3)
	ListHeroNode(head)
	ListHeroNode2(head)

	DeleteHeroNode(head, 2)
	ListHeroNode(head)
}
