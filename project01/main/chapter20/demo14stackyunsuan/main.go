package main

import (
	"errors"
	"fmt"
	"strconv"
)

//Stack 使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int     //表示我们栈最大可存放数个数
	Top    int     //表示栈顶，因为栈顶固定，因为我们直接使用Top
	arr    [20]int //数组模拟
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

//GetTop 获取栈顶元素信息，不出栈
func (s *Stack) GetTop() (val int) {
	if s.Top == -1 {
		//此处报错
		panic(errors.New("stack empty"))
	}
	return s.arr[s.Top]
}

//Empty 是否空栈
func (s *Stack) Empty() bool {
	return s.Top == -1
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

//IsOper 判断一个字符是不是一个运算符[+.-,*,/]
func IsOper(v int) bool {
	if v == 42 || v == 43 || v == 45 || v == 47 {
		return true
	}
	return false
}

//Cal 运算方法
func Cal(num1, num2, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

//Priority 编写一个方法，返回某个运算符的优先级
// [* / => 1, + - => 0]
func Priority(oper int) int {
	if oper == 42 || oper == 47 {
		return 1
	}
	return 0
}

//求出运算： 3+2*6-2
//1.创建两个栈，numStack,operStack
//2.numStack存放数，operStack存放操作符
//3.index:=0
//4.exp计算表达式，一个字符串
//5.如果扫描发现是一个数字，则直接入numstack
//6.如果发现是一个运算法：
//1) 如果operstack是一个空栈，直接入栈
//2) 如果发现operstack栈顶的运算法的优先级大于等于当前准备入栈的运算符的优先级，就从符号栈pop出，并从数栈也pop两个数
//   进行运算，运算后的结果再重新入栈到数栈
//3) 否则，运算符就直接入栈
//7. 如果扫描全部完毕，依次从数栈弹出两个数，从符号位弹出一个符号，计算后入数栈，依此类推。
func main() {
	//数栈
	numStack := &Stack{MaxTop: 20, Top: -1}
	//符号栈
	operStack := &Stack{MaxTop: 20, Top: -1}

	exp := "30+20*6-2/2"
	//定义一个index，帮助扫描exp
	index := 0
	len := len(exp)
	keepNum := ""
	var num1, num2, oper, result int
	for index < len {
		ch := exp[index : index+1] //返回
		//先转成byte切片，然后获取到第一个元素，将byte转成int
		temp := int([]byte(ch)[0])

		if IsOper(temp) { //说明是符号
			if operStack.Empty() { //空栈，直接入栈
				operStack.Push(temp)
			} else {
				//如果发现栈顶的运算法的优先级大于等于当前准备入栈的运算符的优先级，就从符号栈pop出，并从数栈也pop两个数，运算后的结果再重新入栈到数栈
				if Priority(operStack.GetTop()) >= Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()

					result = Cal(num1, num2, oper)
					//fmt.Println(num1, num2, oper, result)
					//将计算的结果入栈
					numStack.Push(result)
				}
				//将要入栈的运算符入栈
				operStack.Push(temp)
			}
		} else { //说明是数
			//处理多位数的思路
			//1.定义一个变量 keepNum string，做拼接
			keepNum += ch
			//2. 每次要向index的后一位字符探测一下，看看是不是运算符，然后处理

			if index == len-1 {
				v, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(v))
			} else {
				//向index后一位测试看看是不是运算符
				if IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					v, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(v))
					keepNum = ""
				}
			}
		}
		index++
	}

	//如果扫描完毕，则依次从数栈弹出两个数，从符号位弹出一个符号，计算后入数栈
	for !operStack.Empty() {
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()

		result = Cal(num1, num2, oper)
		numStack.Push(result)
	}

	result, _ = numStack.Pop()

	fmt.Printf("表达式： %v = %d \n", exp, result)
}
