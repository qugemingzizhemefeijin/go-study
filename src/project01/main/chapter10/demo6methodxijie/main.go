package main

import (
	"fmt"
)

//Person ...
type Person struct {
	Name string
}

//为了
func (p Person) sayHello() {
	fmt.Println(p.Name, " say Hello!")
}

//为了提高效率，通常我们方法是与结构体的指针类型绑定
func (p *Person) speak() {
	fmt.Println(p.Name, " speak!")
}

func (p *Person) String() string {
	return "Person [Name:" + p.Name + "]"
}

// func (p Person) String() string {
// 	return "Person Name:" + p.Name
// }

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

//编写一个方法，可以改变i的值
func (i *integer) change() {
	*i = *i + 1
}

func main() {
	//1) 结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式；
	//2) 如程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理；(实际通常是绑定指针，因为可以加快调用速度)
	p := Person{"Amy"}
	//标准的访问方式应该如下写(编译器自动添加了&p了)：
	//(&p).speak()
	p.speak()

	//3) Go中的方法作用在指定的数据类型上，因为自定义类型，都可以有方法，而不仅仅是struct，比如int，float32等都可以有方法；
	var i integer = 10
	i.print()
	i.change()
	i.print()

	//4) 方法的访问范围控制的规则，和函数一样。方法名首字母小写，只能在本包访问，方法首字母大写，可以在本包和其它包访问；
	//5) 如果一个变量实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出；
	fmt.Println(&p) //主要，这里得传递&p指针，不能直接传递p变量，因为String方法定义的时候，是写的是指针
}
