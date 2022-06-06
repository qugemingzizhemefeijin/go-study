package main

import (
	"fmt"
)

//Stu ...
type Stu struct {
	Name string
	Age  int
}

func main() {
	//方式1：
	//在创建结构体变量时，就直接指定字段的值
	var stu1 = Stu{"小明", 19}
	stu2 := Stu{"小明~", 20}

	//在创建结构体变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序，更稳健
	var stu3 = Stu{
		Name: "小光",
		Age:  22,
	}
	stu4 := Stu{
		Name: "小光~",
		Age:  23,
	}

	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu3)
	fmt.Println(stu4)

	//方式2：返回结构体的指针类型
	var stu5 *Stu = &Stu{"小红", 18}
	stu6 := &Stu{
		Age:  16,
		Name: "小红~",
	}
	fmt.Println(*stu5)
	fmt.Println(*stu6)
}
