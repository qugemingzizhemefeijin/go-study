package main

import "fmt"

var n1 = 100
var n2 = 200
var name = "jack"

//批量声明全局变量
var (
	n3 = 100
	n4 = 300
	name2 = "tom"
)

func main(){
	//如何声明全局变量
	fmt.Println("n1=", n1, "n2=" , n2, "name=", name)
	fmt.Println("n3=", n3, "n4=" , n4, "name2=", name2)
}