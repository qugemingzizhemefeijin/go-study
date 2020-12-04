package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A Say OK", a.Name)
}

func (a *A) hello() {
	fmt.Println("A Hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B Say OK", b.Name)
}

type C struct {
	A
	B
}

type D struct {
	a A
}

type E struct {
	A
	brand string
}

type F struct {
	*A
	brand string
}

type Goods struct {
	Name  string
	price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type DV struct {
	Brand
	int
}

func main() {
	//1. 结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大写或者小写的字段、方法、都可以使用
	// b := &B{}
	// b.A.Name = "tom"
	// b.A.age = 19
	// b.A.SayOk()
	// b.A.hello()
	//2.匿名结构体字段访问可以简化(编译器会先看b对应的类型有没有Name，如果有，则直接调用B类型的Name字段，如果没有则看B中嵌入的匿名结构体A中有没有)
	// b.Name = "Mary"
	// b.age = 29
	// b.SayOk()
	// b.hello()
	//3.当结构体和匿名结构体有先沟通的字段或者方法时，编译器采用就近访问原来访问，如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
	//b := &B{}
	//b.Name = "tom" //OK
	//b.A.Name = "mary" //肯定是给A的匿名结构体赋值
	//b.age = 19 //OK
	//b.SayOk()  //B say ok
	//b.hello()  //A hello ""
	//4.结构体嵌入两个（或多个匿名结构体，如两个匿名结构体有相同的字段和方法（同时结构体本身没有同名的字段和方法）），
	//  在访问时，就必须明确指定匿名结构体名字，否则编译报错
	//var c C
	//c.Name = "abc" //这里会报错
	//c.A.Name = "mart" //这样子就可以了
	//fmt.Println(c)
	//5.如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或方式时，必须带上结构体的名字
	//var d D
	//d.Name = "jetty" //不能这么访问了，因为里面的A不是一个匿名结构体
	//d.a.Name = "tom"
	//fmt.Println(d)
	//6.嵌套匿名结构体后，也可以在创建结构体变量时，直接指定各个匿名结构体字段的值（还可以嵌入匿名结构体的指针）
	e := E{
		A: A{
			Name: "jetty",
			age:  100,
		},
		brand: "朗逸",
	}

	e2 := E{A{"tomcat", 100}, "奔驰"}
	fmt.Println(e)
	fmt.Println(e2)

	//这里用到了指针匿名结构体
	f := F{
		A: &A{
			Name: "resin",
			age:  100,
		},
		brand: "宝马",
	}
	fmt.Println("f=", f, "a=", *f.A)

	//如果一个结构体包含多个匿名结构体，不包含其他基本类型字段，则在初始化的时候（按照名称初始化），可以省略掉结构体名称，如下：
	tv1 := TV{Goods{"电视机", 9999.99}, Brand{"康佳", "河北"}}
	fmt.Println(tv1)
	tv2 := TV{
		Goods{
			Name:  "电视机",
			price: 9988,
		},
		Brand{
			Name:    "创维",
			Address: "湖南",
		},
	}
	fmt.Println(tv2)
	//7.结构体的匿名字段是基本数据类型，也是可以的
	// 注意：1.如果一个结构体有int类型的匿名字段，就不能有第二个；
	//	     2.如果需要有多个int的字段，则必须给int字段取名字
	dv1 := DV{}
	dv1.Brand.Name = "索尼"
	dv1.Brand.Address = "日本"
	dv1.int = 100
	fmt.Println(dv1)
	//8.如一个struct嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承。
	//  案例上面TV结构体已经有了，不过尽量不要使用 多重继承，避免复杂性
}
