package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(s string) string {
		if strings.HasSuffix(s, suffix) {
			return s
		}
		return s + suffix
	}
}

func main() {
	//1.编写一个函数makeSuffix(suffix string)可以接收一个文件后缀名（比如.jpg），并返回一个闭包
	//2.调用闭包，可以传入一个文件名，如果该文件名没有指定后缀（比如.jpg），则返回文件名.jpg，如果已经有.jpg后缀，则返回原文件名。
	//3.要求使用闭包的方式完成
	//4.strings.HasSuffix

	//a := "Hellp.jpg"
	//fmt.Println(strings.HasSuffix(a, ".jpg"))

	ms := makeSuffix(".jpg")
	fmt.Println("makeSuffix = ", ms("aaa"))
	fmt.Println("makeSuffix = ", ms("chen"))

	//总结
	//1.返回的函数和makeSuffix(suffix string)的suffix变量组合成一个闭包，因为返回的函数引用到suffix这个变量
	//2.我们体会一些闭包的好处，如果使用传统的方法，也可以轻松实现这个功能，但是传统方法需要每次传入后缀名，而闭包因为可以保留上次引用的某个值，我们只需要传入一次，就可以反复使用。
}
