package main

import (
	"fmt"
)

func test(n1 int) {
	n1++
	fmt.Println("n1=", n1)
}

//两数之和
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func calc(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}

//1.载调用一个函数时，会给该函数分配一个栈空间，
//2.每个函数对应的栈中，数据空间是独立的，不会混淆
//3.当一个函数执行完毕后，栈空间将被销毁
//4.如果返回多个值时，在接收时，希望忽略某个返回值，则使用_符号表示占位忽略
//5.如果返回值只有一个，返回值类型泪飙可以不写()
func main() {
	test(10)

	sum := getSum(10, 20)
	fmt.Println("sum = ", sum)

	res1, res2 := calc(30, 20)
	fmt.Println("sum=", res1, ",sub=", res2)

	//希望忽略某个返回值，则使用 _ 符号表示占位忽略
	_, res3 := calc(1, 2)
	fmt.Println("res3=", res3)
}
