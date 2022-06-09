package main

import (
	"fmt"
)

//需要保存个数不确定的数据，则需要使用Slice，可以理解成动态数组
//切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制。
//切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度都一样
//切片的长度是可以变化的，因为切片是一个可以动态变化的数组
//切片定义的基本语法：
//var 变量名 []类型，比如：var a []int
func main() {
	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//声明/定义一个切片
	//1. slice就是切片名称
	//2. intArr[1:3]表示slice 引用到intArr这个数组的第2个元素到第3个元素（右侧是开区间）
	slice := intArr[1:3]
	fmt.Println("intArr = ", intArr)
	fmt.Println("slice 的元素是 =", slice)
	fmt.Println("slice 的元素个数 =", len(slice))
	fmt.Println("slice 的容量 =", cap(slice)) //切片的容量是可以动态变化的 [一般为当前初始元素的两倍]

	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	//slice元素有3个，1是指向数组的地址，2是当前长度，3是容量大小
	fmt.Printf("slice[0]的地址=%p slice[0]=%v\n", &slice[0], slice[0])

	//总结：
	//1.slice的确是一个引用类型
	//2.slice从底层来说，其实就是一个数据结构（struct结构体）
	//type slice struct {
	// ptr *[2]int
	// len int
	// cap int
	//}

	slice[1] = 34
	fmt.Println("修改后 intArr = ", intArr)
	fmt.Println("修改后 slice 的元素是 =", slice)

	//第二种创建切片的方式，对于切片，必须make使用
	var s []float64 = make([]float64, 5, 10)
	s[1] = 10
	s[4] = 40
	fmt.Println(s)
	fmt.Println("slice的size=", len(s))
	fmt.Println("slice的cap=", cap(s))

	//第三种，在定义切片时，直接就指定具体数据，使用原理跟make相似，这里cap=3，是因为这个切片给予赋值了初始值
	var v []int = []int{1, 2, 3}
	fmt.Println(v)
	fmt.Println("v的size=", len(v))
	fmt.Println("v的cap=", cap(v))

	//切片遍历================================================================
	//1.for循环常规方式
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice = arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}
	fmt.Println("==============================")
	//2.for range方式遍历切片
	for i, val := range slice {
		fmt.Printf("slice[%v]=%v\n", i, val)
	}
}
