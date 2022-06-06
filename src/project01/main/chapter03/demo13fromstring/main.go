package main

import (
	"fmt"
	"strconv"
)

//string类型转基本类型
func main() {
	//注意事项：
	//要确保string类型能够转成有效的数据，比如把"hello"转成一个整数，不会报错，但是会直接返回0

	var s1, s2, s3 string = "10", "bool", "3.1415926"
	fmt.Println("s1=", s1, ", s2=", s2, ", s3=", s3)

	var a int64
	var b bool
	var c float64

	a, _ = strconv.ParseInt(s1, 10, 64)
	fmt.Printf("a type is %T, a = %d\n", a, a)

	//此函数会返回两个值，,_表示忽略接收第二个值
	b, _ = strconv.ParseBool(s2)
	fmt.Printf("b Type is %T, b = %t\n", b, b)

	c, _ = strconv.ParseFloat(s3, 64)
	fmt.Printf("c Type is %T, c = %f\n", c, c)

	//注意
	var str4 string = "hello"
	var n3 int64 = 11
	n3, _ = strconv.ParseInt(str4, 10, 64)
	//n3没有转换成功，则n3=0
	fmt.Printf("n3 Type is %T, n3 = %v\n", n3, n3)
}
