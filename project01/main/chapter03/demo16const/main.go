package main

import (
	"fmt"
)

func main() {
	//常量在申明的时候必须要赋值，不能使用:=来赋值
	const a = 10
	//常量不能修改
	//a = 20
	//常量只能修饰bool、数值类型(int、float类型)、string类型
	//语法：const identifier [type] = value

	//简洁的写法
	const (
		a1 = 1
		b1 = 2
	)
	fmt.Println(a1, b1)

	//专业的写法
	const (
		d = iota //表示给a赋值为0，b在a的基础上+1，c在b的基础上+1
		e
		f
	)
	fmt.Println(d, e, f)

	const (
		a2     = iota       //默认0
		b2     = iota       //a + 1
		c2, d2 = iota, iota //b+1,b+1
	)
	fmt.Println(a2, b2, c2, d2) //输出0,1,2,2，可以理解成一行递增一次

	//Go中没有常量名必须字母大写的规范
	//仍然通过首字母的小写来控制常量的访问控制

	fmt.Println("exit success")
}
