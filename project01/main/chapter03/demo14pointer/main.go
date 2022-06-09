package main

import (
	"fmt"
)

// 演示golang中指针类型
// 指针的细节说明
// 1.值类型，都有对应的指针类型，形式为*数据类型，蔽日int的对应的指针就是*int,float32对应的指针类型是*float32等
// 2.值类型包括：基本数据类型(int系列，float系列,bool,string)、数组和结构体struct
// 3.引用类型：指针、slice切片、map、管道chan、interface等都是引用类型

// 值类型：变量直接存储值，内存通常在栈中分配
// 引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据，内存通常在堆上分配，当没有任何变量引用这个地址时，
//          该地址对应的数据空间就成为一个垃圾，由GC来回收。
func main() {
	//基本数据类型在内存布局
	// var i int = 10
	// i 的地址是什么？
	// fmt.Println("i的地址=", &i)

	//将i的地址赋值给ptr
	// 1. ptr是一个指针变量
	// 2. ptr的类型是*int
	// 3. ptr 本身的值&i
	// var ptr *int = &i
	// fmt.Println("ptr的地址=", ptr)
	// fmt.Println("ptr本身的地址=", &ptr)
	// fmt.Printf("ptr指向的值=%v\n", *ptr)

	//1.写一个程序，获取一个int变量num的地址，并显示到终端
	var num int = 9
	fmt.Printf("num address = %v\n", &num)
	//2.将num的地址赋给指针ptr，并通过ptr去修改num的值。
	var ptr *int = &num
	*ptr = 10
	fmt.Printf("num value = %d\n", num)

	var a int = 300
	var b int = 400
	var p *int = &a
	*p = 100
	p = &b
	*p = 200
	fmt.Printf("a=%d,b=%d,*ptr=%d", a, b, *p)
}
