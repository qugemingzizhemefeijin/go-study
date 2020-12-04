package main

import (
	"fmt"
	"math"
)

func sushu(n int) {
label1:
	for i := 2; i <= n; i++ {
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				continue label1
			}
		}
		fmt.Println(i)
	}
}

func main() {
	sushu(100)
}
