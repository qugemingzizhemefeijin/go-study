package main

import (
	"fmt"
	"os"
)

//Emp 定义员工结构体
type Emp struct {
	ID   int
	Name string
	Next *Emp
}

//EmpLink 结构体
//我们这里的EmpLink不带表头，即第一个节点就存放雇员
type EmpLink struct {
	Head *Emp
}

//Insert 添加员工的方法，保证添加时，编号从小到大
func (link *EmpLink) Insert(emp *Emp) {
	cur := link.Head   //这个是辅助指针
	var pre *Emp = nil //这是一个辅助指针，在cur前面

	//如果当前EmpLink本身就是空链表
	if cur == nil {
		link.Head = emp
		return
	}
	//如果不是一个空链表，给emp找到对应的位置并插入
	//思路是，让cur和emp比较，让pre始终保持在cur的前面
	for {
		if cur != nil {
			if cur.ID >= emp.ID {
				//找到位置了
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}

	//如果pre==nil代表应该插入到头部
	if pre == nil {
		emp.Next = cur
		link.Head = emp
	} else {
		//直接将pre的Next指向当前，emp的Next指向cur
		pre.Next = emp
		emp.Next = cur
	}
}

//ShowLink 显示当前链表的信息
func (link *EmpLink) ShowLink(no int) {
	if link.Head == nil {
		fmt.Printf("链表%d数据为空\n", no)
		return
	}

	//遍历当前链表，并显示数据
	cur := link.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员ID=%d 名字=%s ->", no, cur.ID, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//根据ID查找雇员
func (link *EmpLink) find(id int) *Emp {
	if link.Head == nil {
		return nil
	}

	cur := link.Head
	for {
		if cur != nil {
			if cur.ID == id {
				return cur
			}
			cur = cur.Next
		} else {
			break
		}
	}
	return nil
}

//Delete 删除指定的雇员
func (link *EmpLink) Delete(id int) bool {
	if link.Head == nil {
		return false
	}
	cur := link.Head
	var pre *Emp = nil

	for {
		if cur != nil {
			if cur.ID == id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			return false
		}
	}

	fmt.Println("=====pre=", pre)
	fmt.Println("=====cur=", cur)
	if pre == nil { //代表是头指针
		link.Head = cur.Next
	} else {
		pre.Next = cur.Next
	}

	return true
}

//HashTable 含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//Insert 给HashTable 编写Insert雇员的方法
func (t *HashTable) Insert(emp *Emp) {
	//使用散列函数确定将该雇员添加到哪个链表
	linkNum := t.hashFun(emp.ID)
	//调用对应的EmpLink进行添加
	t.LinkArr[linkNum].Insert(emp)
}

//Show 显示所有雇员
func (t *HashTable) Show() {
	for i := 0; i < len(t.LinkArr); i++ {
		t.LinkArr[i].ShowLink(i)
	}
}

//编写用于一个散列的方法
func (t *HashTable) hashFun(id int) int {
	return id % 7
}

//编写方法，完成查找
func (t *HashTable) find(id int) *Emp {
	linkNum := t.hashFun(id)
	//调用对应的EmpLink进行查找
	return t.LinkArr[linkNum].find(id)
}

//Delete 编写方法，完成删除操作
func (t *HashTable) Delete(id int) bool {
	linkNum := t.hashFun(id)
	return t.LinkArr[linkNum].Delete(id)
}

//在一个公司，当有新的员工来报道时，要求将该员工的信息加入(ID，姓名)，当输入该员工的ID时，要求查找到员工的所有信息
func main() {
	key := ""
	id := 0
	name := ""

	var hash HashTable

	for {
		fmt.Println("----------------------雇员系统菜单----------------")
		fmt.Println("input 表示添加雇员")
		fmt.Println("delete 表示删除雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 表示退出系统")

		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)

		switch key {
		case "input":
			fmt.Println("输入雇员ID")
			fmt.Scanln(&id)
			fmt.Println("输入雇员名字")
			fmt.Scanln(&name)

			emp := &Emp{
				ID:   id,
				Name: name,
			}
			hash.Insert(emp)
		case "delete":
			fmt.Println("输入要删除的雇员ID")
			fmt.Scanln(&id)
			if hash.Delete(id) {
				fmt.Println("删除成功")
			} else {
				fmt.Println("删除失败")
			}
		case "show":
			hash.Show()
		case "find":
			fmt.Println("输入雇员ID")
			fmt.Scanln(&id)
			emp := hash.find(id)
			if emp == nil {
				fmt.Println("没有找到ID=", id, "的雇员")
			} else {
				fmt.Printf("雇员ID=%d,名字=%s\n", emp.ID, emp.Name)
			}
		case "exit":
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}
}
