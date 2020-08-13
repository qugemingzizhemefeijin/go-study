package main

import (
	"fmt"
)

//AddUpper 累加器
func AddUpper() func(int) int {
	var n int = 10
	var str string = "hello"
	return func(x int) int {
		n = n + x
		str += fmt.Sprintf("%d", x)
		fmt.Println("str=", str)
		return n
	}
}

func main() {
	//闭包就是一个函数和与其相关的引用环境组合的一个整体
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
}
