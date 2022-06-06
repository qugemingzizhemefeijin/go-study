package main

import (
	"fmt"
	"project01/main/chapter06/demo7init/utils" //被引入的包的init先执行
)

var age int = test()

//为了看到全局变量是先被初始化的，我们这里先写函数
func test() int {
	fmt.Println("test()") //1
	return 20
}

//init函数，通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("main init...") //2
}

//init函数介绍
func main() {
	//每一个源文件中都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用，也就是说init会在main函数前被调用
	fmt.Println("main() ... ") //3

	//init函数的细节：
	//1.如果一个文件同时包含全局变量定义,init函数和main函数，则执行的流程是变量定义->init函数->main函数
	//2.init函数最主要的作用，就是完成一些初始化的工作，比如下面案例：
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
	//3.如果main.go和utils.go都含有变量定义，init函数执行的流程又是怎么样的呢？
	// 先执行utils.go的变量定义，init函数，再执行main.go的变量定义，init函数，main函数
}
