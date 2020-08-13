package main

import (
	"fmt"
)

func main() {
	//打印九九乘法表
	var top int = 20
	for i := 1; i < top; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println("")
	}
}
