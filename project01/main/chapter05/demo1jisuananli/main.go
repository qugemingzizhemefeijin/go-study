package main

import (
	"fmt"
)

func main() {
	//编写程序，声明2个int32型变量并赋值，判断两数之和，如果大于等于50，打印“hello world!”

	var i, j int = 30, 20
	if i+j >= 50 {
		fmt.Println("Hello world!")
	}

	//声明两个float64类型并赋值，判断第一个数大于10.0，且第二个数小于20.0，打印两数之和
	var n3, n4 float64 = 11.0, 17.0
	if n3 > 10.0 && n4 < 20.0 {
		fmt.Printf("n3=%v,n4=%v,n3+n4=%v\n", n3, n4, n3+n4)
	}

	// 定义两个变量int32，判断二者的和，是否能被3又能被5整除，打印提示信息。
	var n5, n6 int32 = 30, 15
	n7 := n5 + n6
	if n7%3 == 0 && n7%5 == 0 {
		fmt.Printf("n5=%v,n6%v,n5+n6=%v能被3又能被5整除\n", n5, n6, n7)
	}

	// 判断一个年份是否是闰年，闰年的条件是符合下面二者之一：1年份能被4整除，但不能被100整除。2能被400整除
	var year = 2000
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Printf("年份%v是闰年", year)
	}
}
