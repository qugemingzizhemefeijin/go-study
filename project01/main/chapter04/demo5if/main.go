package main

import "fmt"

func main() {
	//1. 求两个数的最大值
	var i, j int = 10, 8
	var max int
	if i > j {
		max = i
	} else {
		max = j
	}
	fmt.Printf("i=%d,j=%d,最大值为:%d\n", i, j, max)

	//2. 求三个数的最大值
	var m int = 12
	if i > j {
		max = i
	} else {
		max = j
	}
	if max < m {
		max = m
	}

	fmt.Printf("i=%d,j=%d,m=%d,最大值为:%d\n", i, j, m, max)
}
