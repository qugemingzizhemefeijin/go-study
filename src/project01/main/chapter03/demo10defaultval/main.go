package main

import (
	"fmt"
)

func main() {
	//GO汇总所有的数据类型都有默认值
	//整型 0
	//浮点型 0
	//字符串 ""
	//布尔 flase
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	// %v 按照变量的原始值输出
	fmt.Printf("d = %v", d)
}
