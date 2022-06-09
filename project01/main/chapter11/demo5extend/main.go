package main

import (
	"fmt"
)

//Student 学生
type Student struct {
	Name  string
	Age   int
	Score int
}

//将Pupile 和 Graduate 共有的方法也绑定到*Student

//显示他的成绩
func (p *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", p.Name, p.Age, p.Score)
}

//设置分数
func (p *Student) SetScore(score int) {
	//业务判断
	p.Score = score
}

//Pupile 小学生
type Pupile struct {
	Student //嵌入了Student匿名结构体
}

//保留特有的方法
func (p *Pupile) testing() {
	fmt.Println("小学生正在考试中")
}

//Graduate 大学生
type Grduate struct {
	Student //嵌入了Student匿名结构体
}

//保留特有的方法
func (p *Grduate) testing() {
	fmt.Println("大学生正在考试中")
}

func main() {
	//这种调用方式比较麻烦......还有另一种调用方式。。。在后面
	//当我们对结构体嵌入了匿名结构体后，使用的方法会发生变化
	pupil := &Pupile{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()

	//大学生
	grduate := &Grduate{}
	grduate.Student.Name = "Mary"
	grduate.Student.Age = 28
	grduate.testing()
	grduate.Student.SetScore(90)
	grduate.Student.ShowInfo()
}
