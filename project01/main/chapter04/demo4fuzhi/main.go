package main

import "fmt"

func main() {
	//赋值运算符的使用演示
	a := 9
	b := 2
	fmt.Printf("a=%d, b = %d\n", a, b)

	//第一种交换方式
	t := a
	a = b
	b = t
	fmt.Printf("第一次交换：a=%d, b = %d\n", a, b)

	//第二种交换方式
	a += b
	b = a - b
	a = a - b
	fmt.Printf("第二次交换：a=%d, b = %d\n", a, b)

	//第三种交换方式
	a, b = b, a
	fmt.Printf("第三次交换：a=%d, b = %d\n", a, b)
}
