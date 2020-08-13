package main

import (
	"fmt"
)

var (
	//fun1 就是一个全局匿名函数
	fun1 = func(n1, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//如果某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用
	//匿名函数使用方式1：在定义匿名函数时直接调用

	//案例演示：求两个数的和，使用匿名函数的方式完成
	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	//匿名函数使用方式2：将匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数
	//sub的数据类型就是函数类型了。
	sub := func(n1, n2 int) int {
		return n1 - n2
	}
	fmt.Println("sub=", sub(20, 10))

	//匿名函数使用方式3：全局的匿名函数
	res4 := fun1(4, 9)
	fmt.Println("res4=", res4)

	//还能这么定义，牛掰
	var (
		a = 3
	)
	a = 4
	fmt.Println("a=", a)
}
