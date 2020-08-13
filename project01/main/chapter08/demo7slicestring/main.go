package main

import (
	"fmt"
)

func main() {
	//1.string底层是一个byte数组，因为string也可以进行切片处理
	//2.string是不可变的，也就是说不能通过str[0] = 'z'来修改字符串
	//3.如果需要修改字符串，可以先将string => []byte ->修改 ->重写转成string
	var str = "hello golang"
	fmt.Println(str)

	var slice = str[6:]
	fmt.Println("slice=", slice)

	//str[0] = 'z'	//此处会报错的，因为string不可更改的，虽然他底层是byte数组

	//只能通过下列方式改变
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str = ", str)

	//上面的方式可以处理英文和数字，如果有中文，则无法处理，因为一个中文占3-4个字节，因此会出现乱码问题
	//解决办法是 将 string转成 []rune即可，因为[]rune是按照字符来处理的，兼容汉字
	arr2 := []rune(str)
	arr2[0] = '中'
	str = string(arr2)
	fmt.Println("str = ", str)

	//转成切片行不行？比如获取 0,2看看中文能否正常打印出来。【打印出乱码，那如果非要获取前两个怎么办？】
	slice = str[0:2]
	fmt.Println("slice = ", slice)

	//可以先转成rune，再获取切片
	arr3 := []rune(str)
	slice3 := arr3[0:2]
	fmt.Println("slice = ", string(slice3))
}
