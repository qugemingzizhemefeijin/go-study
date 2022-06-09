package main

import (
	"fmt"
)

//InsertSort ...
func InsertSort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		//完成第一次，给第二个元素找到合适的位置并插入
		insertVal := arr[i]
		insertIndex := i - 1

		//从大到小排序，如果当前指定的位置比比较的值要小，则需要将数据后移
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后%v \n", i, arr)
	}
}

//插入排序思想：
//把n个排序的元素看成为一个有序表和一个无序表，开始时有序表中只包含一个元素，无序表中包含有n-1个元素
//排序过程中每次从无序表中取出第一个元素，把它的排序码依次与有序表元素的排序码比较，将它插入到有序表
//适当的位置，使之成为新的有序表
func main() {
	arr := [5]int{23, 0, 12, 56, 34}
	fmt.Println(arr)
	InsertSort(&arr)
	fmt.Println(arr)

	fmt.Println("exit success")
}
