package main

import (
	"fmt"
)

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()")
}

type CInterface interface {
	AInterface
	BInterface
	Call()
}

//如果需要实现CInterface，就需要将AInterface和BInterface的方法都实现
type Human struct {
}

func (h Human) Hello() {
	fmt.Println("Human Hello()")
}

func (h Human) Say() {
	fmt.Println("Human Say()")
}

func (h Human) Call() {
	fmt.Println("Human Call()")
}

type T interface {
}

type Usb interface {
	Say()
}

type Man struct {
}

func (this *Man) Say() {
	fmt.Println("Man Say()")
}

func main() {
	//1) 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量
	//直接使用结构体是会报错的
	// var a1 AInterface
	// a1.Say()

	//可以如下使用
	var stu Stu
	var a AInterface = stu
	a.Say()

	//2.接口中所有的方法都没有方法体，即都是没有实现的方法
	//3.在Go中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口
	//4.一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋给接口类型
	//5.只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
	var i integer = 10
	var b AInterface = i
	b.Say()
	//6.一个自定义类型可以实现多个接口
	//Monster实现了AInterface和BInterface
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()
	//7.Go接口中不能有任何变量
	//8.一个接口(比如A接口)可以继承多个别的接口(B,C接口)，这时如果要实现A接口，也必须将B,C接口的方法也全部实现。
	var human Human
	var c CInterface = human
	c.Call()
	//9.interface类型默认是一个指针(引用类型)，如果没有对interface初始化就使用，那么会输出nil
	//10.孔接口interface{}没有任何方法，所以所有类型都实现了空接口
	var t T = human //OK
	fmt.Println(t)
	var t2 interface{} = human
	fmt.Println(t2)
	t2 = 8.8
	fmt.Println(t2)
	//11.一个接口继承多个接口，则不能有相同的方法名，否则将编译错误
	//12.如果是指针实现的接口方法，则必须用指针来赋值，否则将报错
	var man Man = Man{}
	//var u Usb = man	//Man does not implement Usb
	var u Usb = &man //这样子才不会报错
	u.Say()
	fmt.Println("here", u)
}
