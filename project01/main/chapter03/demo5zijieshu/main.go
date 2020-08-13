package main

// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)

func main(){
	//int8 = 1字节
	//int16 = 2字节
	//int32 = 4字节
	//int64 = 8字节
	//顺带还有无符号类型 uint8 ...
	var i int8 = -128
	fmt.Println("i=", i)

	//int = 32位系统=4字节 int32，64位等于int64
	//uint 上同，只是表示无符号类型

	//rune 有符号，与int32一样
	//byte 无符号，与uint8等价

	//如何去看一个变量的数据类型
	var n1 = 100
	fmt.Printf("n1 的 类型 %T\n", n1)

	//如果在程序查看某个变量的占用字节大小和数据类型（使用较多）
	var n2 int64 = 10
	fmt.Printf("n2 的 类型 %T n2占用的字节数是 %d\n", n2, unsafe.Sizeof(n2))
}