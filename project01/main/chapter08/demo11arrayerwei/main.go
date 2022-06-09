package main

import (
	"fmt"
)

//二维数组
func main() {
	var arr [4][5]int
	arr[0] = [...]int{0, 0, 0, 0, 0}
	arr[1] = [...]int{0, 0, 1, 0, 0}
	arr[2] = [...]int{0, 2, 0, 3, 0}
	arr[3] = [...]int{0, 0, 0, 0, 0}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Printf("arr[0] address = %p\n", &arr[0])
	fmt.Printf("arr[1] address = %p\n", &arr[1])

	fmt.Printf("arr[0][0] address = %p\n", &arr[0][0])
	fmt.Printf("arr[1][0] address = %p\n", &arr[1][0])

	//初始化的几种方式
	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2)

	var arr3 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr3)

	var arr4 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr4)

	arr5 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr5)

	//二维数组遍历,双层for或者for range
	for i := 0; i < len(arr5); i++ {
		for j := 0; j < len(arr5[i]); j++ {
			fmt.Printf("%v\t", arr5[i][j])
		}
		fmt.Println()
	}
	fmt.Println("========================")
	for _, v := range arr5 {
		for _, t := range v {
			fmt.Printf("%v\t", t)
		}
		fmt.Println()
	}
}
