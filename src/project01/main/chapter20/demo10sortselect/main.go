package main

import "fmt"

//SelectSort 编写函数完成排序
func SelectSort(arr *[5]int) {
	//arr[1] 等价 (*arr)[1]

	for j := 0; j < len(arr)-1; j++ {
		//1.先完成将第一个最大值和 arr[0]交换
		//将设 arr[0] 是最大值
		max := arr[j]
		maxIndex := j

		//2.遍历后面1-len(arr) - 1比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}

		//此处交换
		if maxIndex != j {
			arr[maxIndex], arr[j] = arr[j], arr[maxIndex]
		}

		fmt.Printf("第%d次 %v\n", j+1, *arr)
	}
}

//选择排序应用实例：
//有一群牛，颜值分别是 10分、34分、19分、100分、80分，请使用选择排序从高到底进行排序
func main() {
	//定义数据
	array := [5]int{10, 34, 19, 100, 80}
	fmt.Println(array)
	SelectSort(&array)
	fmt.Println(array)
}
