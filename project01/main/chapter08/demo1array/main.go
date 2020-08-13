package main

import (
	"fmt"
)

//数组可以存放多个统一类型数据。数组也是一种数据类型，在Go中，数组是值类型。
//数组的地址就是第一个元素的地址
func main() {
	//1.定义一个数组
	var hens [7]float64
	//2.给数组的每个元素复制
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	hens[6] = 150.0

	//3.遍历数组，求出总体重
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	fmt.Printf("totalWeight=%vKG, avgWeight=%.2fKG\n", totalWeight, totalWeight/float64(len(hens)))
}
