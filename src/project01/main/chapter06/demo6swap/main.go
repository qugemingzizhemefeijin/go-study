package main

import (
	"fmt"
)

func swap(n1 *int, n2 *int) {
	//第一种
	// var t = *n1
	// *n1 = *n2
	// *n2 = t
	//第二种
	//*n1, *n2 = *n2, *n1
	//第三种
	*n1 = *n1 + *n2
	*n2 = *n1 - *n2
	*n1 = *n1 - *n2
}

//练习
func main() {
	//1.编写一个函数swap(n1 *int, n2 *int) 可以交换n1和n2的值
	var n1, n2 int = 10, 20
	fmt.Println("n1=", n1, ",n2=", n2)
	swap(&n1, &n2)
	fmt.Println("n1=", n1, ",n2=", n2)
}
