package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

//类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
func main() {
	var a interface{}
	var point Point = Point{1, 3}
	a = point //ok
	//如果将a赋给一个Point变量？
	var b Point
	// b = a 不可以
	// b = a.(Point) 可以
	//这个就是类型断言，如果不能转换，此处则会报错
	//表示a是否是指向Point类型的变量，如果是就转成Point类型并赋给b变量
	b = a.(Point)
	fmt.Println(b)

	//类型断言其它案例
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //空接口，可以接收任意类型
	// x => float32 [使用类型断言]
	y := x.(float32)
	fmt.Printf("y的类型是 %T 值是=%v\n", y, y)

	//如何在进行断言时，带上检测机制，如果成功就OK，否则也不要报一个panic
	x, ok := x.(float32)
	if ok {
		fmt.Printf("转换成功 x的类型是 %T 值是=%v\n", x, x)
	} else {
		fmt.Println("转换失败")
	}
	fmt.Println("继续执行下面的代码~~~")
}
