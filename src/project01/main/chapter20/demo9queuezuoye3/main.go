package main

import (
	"fmt"
)

//Boy 小孩子的结构体
type Boy struct {
	No   int
	Next *Boy //指向下一个小朋友
}

//AddBoy 编写一个函数，构成单向的环形链表
// num表示小朋友的个数
// 返回该环形链表的头部指针
func AddBoy(num int) *Boy {

	first := &Boy{}  //空节点
	curBoy := &Boy{} //空节点

	// 判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}

	//循环构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		//构成循环列表，需要一个辅助指针
		// 1. 因为第一个小孩比较特殊
		if i == 1 { //第一个小孩
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

//ShowBoy 显示单向的环形链表[遍历]
func ShowBoy(first *Boy) {
	//如果环形链表为空
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩")
		return
	}

	curBoy := first
	for {
		fmt.Printf("小孩编号=%d ->", curBoy.No)
		if curBoy = curBoy.Next; curBoy == first {
			break
		}
	}
	fmt.Println()
}

//GetCount 获取环形链表的节点数量
func GetCount(first *Boy) int {
	if first.Next == nil {
		return 0
	}

	c := 1
	tail := first
	for tail = tail.Next; tail != first; tail = tail.Next {
		c++
	}
	return c
}

//PlayGame 分析思路
//1. 编写一个函数，PlayGame(first *Boy, startNo int, countNum int) startNo>=1 startNo <=小孩总数
//2. 最后我们使用一个算法，按照要求，在环形链表中留下最后一个人
func PlayGame(first *Boy, startNo int, countNum int) {
	//1. 如果是一个空链表，需要单独处理
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	//startNo <= 小孩的总数，如果大于小孩总数，则需要报错
	c := GetCount(first)
	fmt.Println("环形链表长度为=", c)
	if startNo < 1 || startNo > c {
		fmt.Println("传入的小孩起始位置错误")
		return
	}

	//2. 需要定义一个辅助指针，帮助我们删除小朋友
	tail := first
	//3. 让tail指向环形链表的最后一个小孩，因为tail在删除小朋友时需要用到，永远指向当前节点的前一个节点
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	//4.让first移动到startNo[后面我们删除小孩，就以first为准，first指向谁，就删除谁]
	for i := 1; i < startNo; i++ {
		first = first.Next
		tail = tail.Next
	}

	//5.开始数 countNum，然后就删除first 指向的小孩
	for {
		//开始数countNum - 1 次
		for i := 1; i < countNum; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号%d 出圈 -> \n", first.No)
		//删除first指向的节点[让first指向下一个，tail下一个指向移动后的first]
		first = first.Next
		tail.Next = first

		//判断，如果first = tail ，则代表只有最后一个小孩子了，直接出列，并退出for
		if first == tail {
			break
		}
	}
	fmt.Printf("小孩编号%d 出圈 -> \n", first.No)
}

//约瑟夫问题：设编号1-n个小朋友围坐一圈，约定编号k(1<=k<=n)的人从1开始报数，数到m的那个人出列，它的下一位又从1开始报数
//数到m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列。

//就是使用一个循环链表来处理约瑟夫问题，数到m则删除节点并出列。记录出列顺序即可。
func main() {
	first := AddBoy(50)
	//显示
	ShowBoy(first)
	//出圈
	PlayGame(first, 20, 31)
}
