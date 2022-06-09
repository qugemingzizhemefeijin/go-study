package main

import (
	"fmt"
)

func sum(n1, n2 int) int {
	//当执行到defer时，暂时不执行，会将defer后的语句压入到独立的栈中（defer栈）
	//当函数执行完毕后，再从defer栈，按照先入后厨的方式出栈，执行
	//在defer将语句放入到栈时，也会将相关的值拷贝同时入栈。
	defer fmt.Println("ok1 n1=", n1) //10
	defer fmt.Println("ok2 n2=", n2) //20

	n1++ //11
	n2++ //21

	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

//在函数中，程序员经常要创建资源，为了在函数执行完毕后，及时释放资源，Go的设计者提供了defer（延时机制）。
func main() {
	res := sum(10, 20)
	fmt.Println("main res=", res)
}
