package main

import (
	"encoding/json"
	"fmt"
)

//Point 点
type Point struct {
	x int
	y int
	m string
}

//Rect 矩形
type Rect struct {
	leftUp, rightDown Point
}

//C ...
type C struct {
	Num int
}

//B ...
type B struct {
	Num int
}

//Monster ...
type Monster struct {
	Name  string `json:"name"` //结构体的标签
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	r1 := Rect{Point{1, 2, "a"}, Point{3, 4, "b"}}
	fmt.Println(r1)
	//r1有4个int,在内存中是连续分布的
	//打印地址
	fmt.Printf("r1.leftUp.x 地址=%p, r1.leftUp.y 地址=%p, r1.leftUp.m 地址=%p\n", &r1.leftUp.x, &r1.leftUp.y, &r1.leftUp.m)
	fmt.Printf("r1.rightDown.x 地址=%p, r1.rightDown.y 地址=%p, r1.rightDown.m 地址=%p\n", &r1.rightDown.x, &r1.rightDown.y, &r1.rightDown.m)

	fmt.Println("=========================")
	//存储可变类型，会咋样的？[证明string其实是维护的一个byte数组的指针]
	r1.leftUp.m = "assadasdasdasdadadasdb"
	fmt.Printf("r1.leftUp.x 地址=%p, r1.leftUp.y 地址=%p, r1.leftUp.m 地址=%p\n", &r1.leftUp.x, &r1.leftUp.y, &r1.leftUp.m)
	fmt.Printf("r1.rightDown.x 地址=%p, r1.rightDown.y 地址=%p, r1.rightDown.m 地址=%p\n", &r1.rightDown.x, &r1.rightDown.y, &r1.rightDown.m)

	// 结构体是用户单独定义的类型，和其它类型进行转换时需要完全相同（包括名字/个数和类型）
	var c C
	var b B
	//a = b //? 报错
	c = C(b) //这种可以转换，要求结构体的字段要完全一样，包括名字/个数和类型
	fmt.Println(c, b)

	// 结构体进行type重新定义（相当于取别名），Go认为是新的数据类型，但是相互间可以强转

	// 结构体的每一个字段上，可以写一个tag，该tag可以通过反射机制获取，常见使用场景就是序列化和反序列化
	monster := Monster{"牛魔王", 500, "野蛮冲撞"}
	//json序列化，返回的是[]byte切片
	// json Marshal函数中使用了反射
	jsonMonster, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json错误，err", err)
	} else {
		//如果不特殊处理，这里就尴尬了{"Name":"牛魔王","Age":500,"Skill":"野蛮冲撞"}，属性是大写的。
		//如果把属性改成小写，则json出来返回的是空串，是因为json无法访问到小写属性，这里就需要使用到struct的tag
		fmt.Println(string(jsonMonster))
	}
}
