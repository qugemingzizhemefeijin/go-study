package main

import (
	"fmt"
)

func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) { //这里type是一个关键字，固定写法
		case bool:
			fmt.Printf("第%d个参数是 bool 类型，值是%v\n", i, x)
		case float64:
			fmt.Printf("第%d个参数是 float64 类型，值是%v\n", i, x)
		case float32:
			fmt.Printf("第%d个参数是 float32 类型，值是%v\n", i, x)
		case int, int32, int64:
			fmt.Printf("第%d个参数是 int 类型，值是%v\n", i, x)
		case nil:
			fmt.Printf("第%d个参数是 bool 类型，值是%v\n", i, x)
		case string:
			fmt.Printf("第%d个参数是 string 类型，值是%v\n", i, x)
		case Student:
			fmt.Printf("第%d个参数是 Student 类型，值是%v\n", i, x)
		case *Student:
			fmt.Printf("第%d个参数是 *Student 类型，值是%v\n", i, x)
		default:
			fmt.Printf("第%d个参数是 未知 类型,值是%v\n", i, x)
		}
	}
}

type Student struct {
}

//写一函数，循环判断传入参数类型
func main() {
	a := 1
	b := 1.0
	var c float32 = 3.0
	var d Student
	e := &Student{}
	TypeJudge(a, b, c, d, e)
}
