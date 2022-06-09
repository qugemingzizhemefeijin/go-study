package main

import (
	"fmt"
)

func main() {
	//var a [3]int						// 定义长度为3的int型数组，元素全部为0
	//var b = [...]int{1, 2, 3}   		// 定义长度为3的int型数组，元素为1, 2, 3
	//var c = [...]int{2: 3, 1: 2}		// 定义长度为3的int型数组，元素为0, 2, 3
	//var d = [...]int{1, 2, 4: 5, 6}	// 定义长度为6的int型数组，元素为1, 2, 0, 0, 5, 6
	var a = [...]int{1, 2, 3} //
	var b = &a                // b是指向a数组的指针

	fmt.Println(a[0], a[1]) // 打印数组的前2个元素
	fmt.Println(b[0], b[1]) // 通过数组指针访问数组元素的方式和数组类似

	fmt.Println("====================")

	for i, v := range b { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	fmt.Println("====================")

	// 常见的遍历数组方式
	// （用for range方式迭代的性能可能会更好一些，这种迭代可以保证不会出现数组越界的情形，每轮迭代对数组元素的访问时可以省去对下标越界的判断）
	// 1
	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
	// 2
	for i, v := range a {
		fmt.Printf("a[%d]: %d\n", i, v)
	}
	// 3
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}

	fmt.Println("====================")

	// 用for range方式迭代，还可以忽略迭代时的下标
	var times [5][0]int // 虽然第一维数组有长度，但是数组的元素[0]int大小是0，因此整个数组占用的内存大小依然是0，没有付出额外的内存代价。
	for range times {
		fmt.Println("Hello")
	}

	fmt.Println("====================")

	// 数组不仅仅可以用于数值类型，还可以定义字符串数组、结构体数组、函数数组、接口数组、管道数组等等；

	var s1 = [2]string{"Hello", "World"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好"}

	fmt.Println(s1, s2, s3)

	fmt.Printf("a: %T\n", a)  // 打印数组类型
	fmt.Printf("a: %#v\n", a) // 打印数组详细信息

}
