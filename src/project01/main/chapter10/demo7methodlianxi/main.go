package main

import (
	"fmt"
)

//1. 编写结构体(MethodUtils)，编写一个方法，方法不需要参数，在方法中打印10*8的举行，在main方法中调用该方法

//MethodUtils ...
type MethodUtils struct {
	//Rows, Cols int
}

func (mu *MethodUtils) rect() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//2. 编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形

func (mu *MethodUtils) drawRect(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//3. 编写一个方法算该矩形的面积（可以接收长len，宽width），将其作为方法返回值。在main方法中调用该方法，接收返回的面积值并打印
func (mu *MethodUtils) area(len, width int) int {
	return len * width
}

//4. 编写方法：判断一个数是奇数还是偶数
func (mu *MethodUtils) oadd(n int) {
	if n%2 == 0 {
		fmt.Println(n, "是偶数")
	} else {
		fmt.Println(n, "是奇数")
	}
}

//5. 根据行、列、字符打印对应行数和列数的字符，比如：行：3，列：2，字符*，则打印相应的效果
func (mu *MethodUtils) print3(n int, m int, key string) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

//7. 在MethodUtils结构体编个方法，从键盘接收整数（1-9），打印对应乘法表：
func (mu *MethodUtils) cfTable() {
	fmt.Println("请输入1-9的整数：")
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func (mu *MethodUtils) print4(arr [3][3]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

//8. 编写方法，使给定的一个二维数组（3x3）转置：
func (mu *MethodUtils) rechange() {
	arr := [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	mu.print4(arr)
	fmt.Println("=============")

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr[i]); j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
	mu.print4(arr)
}

//6. 定义小小计算器结构体（Calcuator），实现加减乘除四个功能，实现形式1：分4个方法完成；实现形式2：用一个方法搞定；

//Calcuator ...
type Calcuator struct {
	Num1, Num2 float64
}

func (c *Calcuator) getSum() float64 {
	return c.Num1 + c.Num2
}

func (c *Calcuator) getSub() float64 {
	return c.Num1 - c.Num2
}

func (c *Calcuator) calc(operator byte) float64 {
	switch operator {
	case '+':
		return c.Num1 + c.Num2
	case '-':
		return c.Num1 - c.Num2
	case '*':
		return c.Num1 * c.Num2
	case '/':
		return c.Num1 / c.Num2
	default:
		fmt.Println("输入操作符错误")
		return 0.0
	}
}

func main() {
	m := MethodUtils{}
	m.rect()
	fmt.Println("================")
	m.drawRect(10, 8)
	fmt.Println("================")
	area := m.area(10, 20)
	fmt.Println("面积为：", area)
	fmt.Println("================")
	m.oadd(9)
	m.oadd(20)
	fmt.Println("================")
	m.print3(7, 20, "+")
	fmt.Println("================")
	c := Calcuator{30, 10}
	fmt.Println("sum=", fmt.Sprintf("%.2f", c.getSum()))
	fmt.Println("sub=", fmt.Sprintf("%.2f", c.getSub()))
	fmt.Println("================")
	fmt.Println("+=", fmt.Sprintf("%.2f", c.calc('+')))
	fmt.Println("-=", fmt.Sprintf("%.2f", c.calc('-')))
	fmt.Println("*=", fmt.Sprintf("%.2f", c.calc('*')))
	fmt.Println("/=", fmt.Sprintf("%.2f", c.calc('/')))
	fmt.Println("================")
	m.cfTable()
	fmt.Println("================")
	m.rechange()
}
