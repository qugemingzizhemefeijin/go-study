package main

import "fmt"

func main(){
	//小数类型只有float32/float64 浮点数都是有符号的，一般推荐使用float64，少用float32
	var price float32 = 99.87
	fmt.Println("price=", price)

	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)

	//golang的浮点类型默认是float64位
	var num5 = 1.1
	fmt.Printf("num5的数据类型是%T\n", num5)

	//有两种便是形式
	//1十进制数边是方式
	//2科学计数方式
	num6 := 5.12
	num7 := .123 // = 0.123
	fmt.Println("num6=", num6, "num7=", num7)

	num8 := 5.1234e2
	num9 := 5.1234E2
	num10 := 5.1234E-2
	fmt.Println("num8=", num8, "num9=", num9," num10=", num10)
}