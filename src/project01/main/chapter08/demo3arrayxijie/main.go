package main

import "fmt"

//GO中数组的长度是类型的一部分，必须要写数组长度
func test01(arr [3]int) {
	arr[0] = 88
}

func test02(arr *[3]int) {
	(*arr)[0] = 88
}

func main() {
	//1.数组是多个相同类型数据的组合，一个数组一旦声明/定义，其长度是固定的，不能动态变化
	//2.var arr []int这时arr就是一个slice切片
	//3.数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用。
	//4.数组创建后，如果没有赋值，则有默认值。数值类型默认为0，字符串数组默认为""，bool数组默认为false
	//5.使用数组的步骤 1.声明数组并开辟空间 2给数组各个元素赋值 3使用数组
	//6.数组的下标是从0开始的
	//7.数组下标必须在指定范围内使用，否则报panic：数组越界，比如var arr [5]int 则有效小标为0-4
	//8.GO的数组属于值类型，在默认情况下是值传递，因为会进行值拷贝。数组间不会相互影响
	//9.如想在其它函数中，去修改原来的数组，可以使用引用传递（指针方式）
	//10.长度是数组类型的一部分，在传递函数参数时，需要考虑数组的长度。

	//sss := [...]int{1, 2, 3, 4}
	//fmt.Println(sss)

	// var arr01 [3]float32
	// var arr02 [3]string
	// var arr03 [3]bool
	// fmt.Printf("arr01=%v, arr02=%v, arr03=%v\n", arr01, arr02, arr03)

	//数组是进行值拷贝
	arr := [3]int{11, 22, 33}
	test01(arr)
	fmt.Println("main arr = ", arr)

	//如果要在其他地方修改，则需要传递引用
	test02(&arr)
	fmt.Println("pointer arr = ", arr)
}
