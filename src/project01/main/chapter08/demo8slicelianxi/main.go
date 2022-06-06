package main

import (
	"fmt"
)

//练习
//编写一个函数fbn(n int)，要求完成
//1) 可以接收一个 n int
//2) 能够将斐波那契的数列放到切片中
//3) 提示，斐波那契的数列形式：
//   arr[0] = 1;arr[1] = '1'; arr[2] = 2;arr[3] = 3;arr[4] = 5; arr[5] = 8

func fbn(n int) []int {
	if n <= 0 {
		fmt.Println("传入参数有误")
		return nil
	}
	var arr []int = make([]int, n, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			arr[i] = 1
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return arr
}

func main() {
	arr := fbn(10)
	fmt.Println("arr=", arr)
}
