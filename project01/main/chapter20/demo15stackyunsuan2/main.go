package main

import (
	"errors"
	"fmt"
	"strconv"
)

//Stack 使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int             //表示我们栈最大可存放数个数
	Top    int             //表示栈顶，因为栈顶固定，因为我们直接使用Top
	arr    [20]interface{} //数组模拟
}

//Push ...
func (s *Stack) Push(val interface{}) (err error) {
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
func (s *Stack) Pop() (val interface{}) {
	//先判断栈是否为空
	if s.Top == -1 {
		panic(errors.New("stack empty"))
	}

	val = s.arr[s.Top]
	s.Top--
	return
}

//GetTop 获取栈顶元素信息，不出栈
func (s *Stack) GetTop() (val interface{}) {
	if s.Top == -1 {
		return nil
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
		fmt.Printf("arr[%d] = %v \n", i, s.arr[i])
	}
}

//IsSymbol 判断是否是符号
func IsSymbol(b string) bool {
	return b == "(" || b == ")" || b == "+" || b == "-" || b == "*" || b == "/"
}

//IsPriority 比较符号a和b优先级，如果a优先b则返回true （a表示当前元素,b是栈顶元素）
func IsPriority(a, b string) bool {
	if a == ")" {
		return true
	}
	if b == "(" {
		return false
	}

	//如果a=+或-，则b不管是什么，都得出栈
	if a == "+" || a == "-" {
		return true
	} else if b == "*" || b == "/" { //否则a肯定是 *或/，那么b也要*或/才能出栈
		return true
	}
	return false
}

//TransferSuffixExpression 将中缀表达式转换成后缀表达式栈结构
func TransferSuffixExpression(exp string) *Stack {
	operStack := &Stack{MaxTop: 20, Top: -1}
	symbolStack := &Stack{MaxTop: 20, Top: -1}

	index := 0
	len := len(exp)
	stringNum := ""
	//此处依次解析exp
	for index < len {
		ch := exp[index : index+1]
		if IsSymbol(ch) { //如果是符号，则判断符号栈顶的优先级
			for !operStack.Empty() {
				flag := false
				top := operStack.GetTop().(string)
				//如果当前符号是(或者低于栈顶元素，则栈顶元素出栈到符号栈中
				if IsPriority(ch, top) {
					s := operStack.Pop()
					//弹出栈顶元素，并加入到symbolStack
					if s.(string) != "(" {
						symbolStack.Push(s)
						flag = true
					}
				}
				//控制跳出此循环
				if !flag {
					break
				}
			}
			//fmt.Println("=======")
			//operStack.List()
			//fmt.Println("=======")
			if ch != ")" {
				operStack.Push(ch)
			}
		} else {
			//此处代表是数字，则需要继续探查是否下面一个符号还是数字
			stringNum = stringNum + ch

			if index+1 >= len {
				symbolStack.Push(stringNum)
				break
			} else {
				b := []byte(exp[index+1 : index+2])[0]
				//如果不是的话，则将入到符号栈中，并将stringNum重置
				if b < 48 || b > 57 {
					symbolStack.Push(stringNum)
					stringNum = ""
				}
			}
		}
		index++
	}

	//最后判断operStack是否还有操作符，如果有的话，则整体全部入符号栈
	for !operStack.Empty() {
		s := operStack.Pop()
		str := s.(string)
		symbolStack.Push(str)
	}

	return symbolStack
}

//Cal 两数计算
func Cal(num1, num2 int, oper string) int {
	res := 0
	switch oper {
	case "+":
		res = num2 + num1
	case "-":
		res = num2 - num1
	case "*":
		res = num2 * num1
	case "/":
		res = num2 / num1
	default:
		fmt.Println("===操作符=====异常")
	}
	return res
}

//CalResult 将后缀表达式栈计算出结果
//规则：从左到右遍历表达式的每个数字和符号，遇到是数字就进栈，遇到符号，就将处于栈顶两个数字出栈，进行计算，运算结果进栈，一直到最终获得结果
func CalResult(symbolStack *Stack) int {
	stack := &Stack{MaxTop: 20, Top: -1}

	//将数组转换成切片
	var arr []interface{} = symbolStack.arr[:symbolStack.Top+1]
	str := ""

	for i := 0; i < len(arr); i++ {
		str = arr[i].(string)

		if IsSymbol(str) { //如果遇到符号，则弹出两个数字出栈
			num1 := stack.Pop().(int)
			num2 := stack.Pop().(int)

			stack.Push(Cal(num1, num2, str))
		} else {
			num, _ := strconv.ParseInt(str, 10, 64)
			stack.Push(int(num))
		}
	}

	return stack.Pop().(int)
}

//运算2
//将9+(3-1)*3+10/2的结果计算出来
//本身计算式是 中缀表达式，我们可以将其转化成后缀表达式 9 3 1 - 3 * + 10 2 / +
//中缀表达式转后缀表达式的规则：
//从左到右遍历中缀表达式的每个数字和符号，若是数字就输出，即成为后缀表达式的一部分；
//若是符号，则判断其与栈顶符号的优先级，是右括号或优先级低于栈顶符号(乘除优先加减)则栈顶元素依次出栈并输出，并将当前符号入栈
//一直到最终输出后缀表达式为止
func main() {
	exp := "9+(30-10)*3+10/2-22/2*3"
	//exp := "22/2*3"
	fmt.Println("exp:", exp)
	symbolStack := TransferSuffixExpression(exp)
	symbolStack.List()

	//最终计算结果为
	res := CalResult(symbolStack)
	fmt.Println("res=", res)

	fmt.Println("exit success")
}
