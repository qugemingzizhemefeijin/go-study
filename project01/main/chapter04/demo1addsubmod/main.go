package main

import (
	"fmt"
)

// 1. 对于除号，他的整数除和小数除是有区别的：整数之间坐除法时，将舍去小数部分
// 2. 当对一个数取模时，可以等价 a % b = a - a / b * b
// 3. Golang的自增自减只能当做一个独立语言使用，不能这样使用b:= a++ 或者 b := a--，并且 ++和--只能在变量的后面，不能有--a、++a的使用。
// 4.
func main() {
	//如果运算的都是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println(10 / 4)
	var n1 float32 = 10 / 4
	fmt.Println(n1)
	//如果我们希望保留小数部分，则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	//演示 % 的使用
	// 看公式 a % b = a - a / b * b
	fmt.Println("10 % 3 = ", 10%3)
	fmt.Println("-10 % 3 = ", -10%3)
	fmt.Println("10 % -3 = ", 10%-3)
	fmt.Println("-10 % -3 = ", -10%-3)

	// ++ -- 的使用
	var n3 int = 10
	n3++
	fmt.Println("i=", n3)
	n3--
	fmt.Println("i=", n3)
}
