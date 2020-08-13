package main

import (
	"fmt"
)

func main() {
	//1.打印1-100之间所有是9的倍数的证书和个数及总和
	var total, count int
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			total += i
			count++
		}
	}
	fmt.Printf("整数个数:%d,总和:%d\n", total, count)

	//2.完成下面表达式输出
	j := 6
	for i := 0; i <= j; i++ {
		fmt.Printf("%d + %d = %d \n", i, j-i, j)
	}

	//go中没有while和do while
	//模拟while
	var i int = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello,world", i)
		i++
	}
	//for循环模拟do while实现
	i = 1
	for {
		fmt.Println("hello,world", i)
		i++
		if i > 10 {
			break
		}
	}
}
