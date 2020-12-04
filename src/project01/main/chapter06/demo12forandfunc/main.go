package main

import (
	"fmt"
)

//封装函数打印金字塔
func jinzita(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i*2+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

//封装函数打印乘法表
func chengfabiao(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, i*j)
		}
		fmt.Println("")
	}
}

func main() {
	jinzita(5)
	chengfabiao(6)
}
