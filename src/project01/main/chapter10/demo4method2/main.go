package main

import (
	"fmt"
)

//1) 给Person结构体添加speak方法，输出xxx是一个好人
//2) 给Person结构体添加jisuan方法，可以计算从1+..+1000的结果，说明方法体内可以函数一样，进行各种计算
//3) 给Person结构体jisuan2方法，该方法可以接收一个数n，计算从1++..+n的结果
//4) 给Person结构体添加getSum方法，可以计算两个数的和，并返回结果

//Person ...
type Person struct {
	Name string
}

func (person Person) speak() {
	fmt.Println(person.Name, "is a good man")
}

func (person Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(person.Name, "计算的结果是：", res)
}

func (person Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(person.Name, "计算的结果是：", res)
}

func (person Person) getSum(n1, n2 int) (sum int) {
	sum = n1 + n2
	return
}

func main() {
	var p Person
	p.Name = "tom"
	p.speak()
	p.jisuan()
	p.jisuan2(10)
	res := p.getSum(10, 20)
	fmt.Println("res=", res)
}
