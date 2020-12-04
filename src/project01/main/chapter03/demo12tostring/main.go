package main

import (
	"fmt"
	"strconv"
)

//在程序中经常会用string转基本类型，或者基本类型转string
func main() {
	// 基本类型转string类型
	// 1. 方式1 fmt.Sprintf("%参数", 表达式) [推荐使用]
	// 2. 方式2 使用strconv包的函数

	// 方式1
	var num1 int8 = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'

	var str string

	//使用第一种方式转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str = %q, str type %T\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str = %q, str type %T\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str = %q, str type %T\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str = %q, str type %T\n", str, str)

	fmt.Println("==================================")

	// 使用第二种方式转换
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str = %q, str type %T\n", str, str)

	// 'f'代表格式
	// prec 10 表示小数位保留10位
	// 64代表这个小数是 64位
	str = strconv.FormatFloat(float64(num2), 'f', 10, 64)
	fmt.Printf("str = %q, str type %T\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str = %q, str type %T\n", str, str)
}
