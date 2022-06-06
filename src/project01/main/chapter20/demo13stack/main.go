package main

import (
	"errors"
	"fmt"
)

//Stack 使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int    //表示我们栈最大可存放数个数
	Top    int    //表示栈顶，因为栈顶固定，因为我们直接使用Top
	arr    [5]int //数组模拟
}

//Push ...
func (s *Stack) Push(val int) (err error) {
	//先判断栈顶是否满了
	if s.Top == s.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}

	s.Top++
	//放入数据
	s.arr[s.Top] = val
	return
}

//Pop 出栈
func (s *Stack) Pop() (val int, err error) {
	//先判断栈是否为空
	if s.Top == -1 {
		err = errors.New("stack empty")
		return
	}

	val = s.arr[s.Top]
	s.Top--
	return
}

//List 遍历
func (s *Stack) List() {
	//先判断栈是否为空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	//不为空则从栈顶遍历
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d \n", i, s.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,  //表示最多存放5个数到栈中
		Top:    -1, //表示栈顶的位置
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.List()

	v, _ := stack.Pop()
	fmt.Println("v=", v)
	v, _ = stack.Pop()
	fmt.Println("v=", v)
	stack.List()

	fmt.Println("exit success")
}
