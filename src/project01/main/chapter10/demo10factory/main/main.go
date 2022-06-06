package main

import (
	"fmt"
	"go_code/project01/main/chapter10/demo10factory/model"
)

//Go的结构体没有构造函数，通常可以使用工厂模式来解决这个问题
//如果Student的首字母小写，则只能使用工厂模式来解决了，除非S字母大写

func main() {
	// var stu = model.Student{
	// 	Name:  "tom",
	// 	Score: 78.9,
	// }

	//当student结构体首字母小写，我们可以通过工厂模式来解决
	var stu = model.NewStudent("amy", 88.8)
	fmt.Println(stu) //返回的是指针
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
