package main

import (
	"fmt"
)

//GO中方法是作用在指定的数据类型上的（即：和指定的数据类型绑定），因此自定义类型，都可以有方法，而不仅仅是struct

//Person ...
type Person struct {
	Name string
}

//给Person类型绑定一个方法
//person表示哪个Person变量调用，这个p就是它的副本
func (person Person) test() {
	person.Name = "jack" //这里改变了，但是不影响外部的，其实还是值拷贝
	fmt.Println("test() ", person.Name)
}

func main() {
	p := Person{"trump"}
	p.test()
	fmt.Println("main() person.name", p.Name) //输出trump
}
