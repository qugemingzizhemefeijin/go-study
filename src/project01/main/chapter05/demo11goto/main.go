package main

import (
	"fmt"
)

func main() {
	// here:
	// 	for i := 0; i < 2; i++ {
	// 		for j := 1; j < 4; j++ {
	// 			if j == 2 {
	// 				continue here
	// 			}
	// 			fmt.Println("i=", i, "j=", j)
	// 		}
	// 	}
	//1.continue 实现1-100之内的奇数
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 1 {
	// 		fmt.Println(i)
	// 	}
	// }
	//2.从键盘读入个数不确定的整数，并判断读入的正数和负数的个数，输入位0时结束程序
	// for {
	// 	var a int
	// 	fmt.Scanln(&a)
	// 	if a > 0 {
	// 		fmt.Println("a=", a, "是正数")
	// 	} else if a < 0 {
	// 		fmt.Println("a=", a, "是负数")
	// 	} else {
	// 		break
	// 	}
	// }

	// 下面演示goto的使用
	fmt.Println("ok1")
	goto abc
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
abc:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}
