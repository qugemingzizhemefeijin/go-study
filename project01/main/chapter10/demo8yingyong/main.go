package main

import (
	"fmt"
)

//1.编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float64类型
//2.结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值
//3.在main方法中，创建Student结构体变量，并访问say方法，并将调用结果打印输出

//Student ...
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v], gender=[%v], age=[%v], id=[%v], score=[%v]",
		student.name, student.gender, student.age, student.id, student.score)

	return infoStr
}

//1.编程创建一个Box结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终端获取
//2.声明一个方法获取立方体的体积
//3.创建一个Box结构体变量，打印给定尺寸的立方体的体积

//Box ...
type Box struct {
	length, width, height float64
}

func (box *Box) cubage() float64 {
	return box.length * box.width * box.height
}

//1.一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20元，其它情况门票免费。
//2.请编写Visitor结构体，根据年龄段决定能够购买的门票价格并输出

//Visitor ...
type Visitor struct {
	name string
	age  int
}

func (v *Visitor) pay() {
	if v.age > 18 {
		fmt.Printf("%v的年龄为：%v，门票价格为：20元\n", v.name, v.age)
	} else {
		fmt.Printf("%v的年龄为：%v，门票免费\n", v.name, v.age)
	}
}

func main() {
	var stu = Student{
		name:   "Tom",
		gender: "male",
		age:    18,
		id:     20,
		score:  99.9,
	}

	fmt.Println(stu.say())

	fmt.Println("========================")
	var length, width, height float64
	fmt.Println("请输入长宽高：")
	fmt.Scanf("%f %f %f", &length, &width, &height)

	box := Box{length, width, height}
	tiji := box.cubage()

	fmt.Printf("立方体，长=%v,宽=%v,高=%v,体积=%v\n", length, width, height, tiji)

	fmt.Println("========================")
	var visitor Visitor
	fmt.Println("请输入姓名：")
	fmt.Scanln(&visitor.name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&visitor.age)

	visitor.pay()
}
