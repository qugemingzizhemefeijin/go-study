package main

import (
	"fmt"
)

//冒泡排序
//总结：
//1.一共会经过arr.length - 1次的轮数比较，每一轮将会确定一个数的位置
//2.每一轮的比较在主键的减少。
//3.当发现前面的一个数比后面的一个数大的时候，就进行交换
func main() {
	var arr = [...]int{24, 69, 80, 57, 13}
	fmt.Println("arr=", arr)

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("arr=", arr)
}
