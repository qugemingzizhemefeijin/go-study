package main

import (
	"fmt"
)

//Golang设计者为了编程方便，提供了一些函数，这些函数可以直接使用，我们称为Go的内置函数。
func main() {
	//1.len：用来求长度，比如string、array、slice、map、channel
	//2.new：用来分配内存，主要用来分配值类型，比如int、float32、struct...返回的是指针
	num1 := 100
	fmt.Printf("num1的类型%T，num1的值=%v，num1的地址%v \n", num1, num1, &num1)

	num2 := new(int) //*int
	*num2 = 100
	fmt.Printf("num2的类型%T，num2的值=%v，num2的地址%v \nnum2这个指针指向的值=%v", num2, num2, &num2, *num2)
	//3.make：用来分配内存，主要用来分配引用类型，比如chan、map、slice。
}
