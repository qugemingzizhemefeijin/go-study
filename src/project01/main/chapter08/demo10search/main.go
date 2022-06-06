package main

import (
	"fmt"
)

//1.有一个数列：白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王，判断是否包含此名称
func test01() {
	fmt.Println("请输入明教四大护法:")
	var str string
	fmt.Scanln(&str)

	arr := [...]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	for i := 0; i < len(arr); i++ {
		if str == arr[i] {
			fmt.Println("你输入的", str, "包含在数列中，下标为：", i)
			return
		}
	}
	fmt.Println("你输入的", str, "不包含在数列中")
}

//2.请对一个有序数组进行二分查找｛1,9,10,89,1000,1234｝输入一个数看看该数组是否存在此数，并且求出下标，如果没有就提示“没有这个数”
func test02() {
	arr := [...]int{1, 9, 10, 89, 1000, 1234}
	fmt.Println("请输入数字：")
	var i int
	fmt.Scanln(&i)
	testSearch(i, &arr, 0, len(arr)-1)
}

func testSearch(i int, arr *[6]int, startIndex int, endIndex int) {
	if startIndex > endIndex {
		fmt.Println("没有找到")
		return
	}
	idx := (startIndex + endIndex) / 2
	if (*arr)[idx] == i {
		fmt.Println("你输入的", i, "包含在数列中，下标为：", idx)
	} else if (*arr)[idx] > i {
		testSearch(i, arr, startIndex, idx-1)
	} else {
		testSearch(i, arr, idx+1, endIndex)
	}
}

func main() {
	//test01()
	test02()
}
