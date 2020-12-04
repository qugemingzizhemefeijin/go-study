package main

import (
	"fmt"
)

//编写一个学生考试系统

//小学生
type Pupil struct {
	Name  string
	Age   int
	Score int
}

//显示他的成绩
func (p *Pupil) showInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", p.Name, p.Age, p.Score)
}

//设置分数
func (p *Pupil) SetScore(score int) {
	//业务判断
	p.Score = score
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中")
}

//代码冗余了
//大学生
type Graduate struct {
	Name  string
	Age   int
	Score int
}

//显示他的成绩
func (p *Graduate) showInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", p.Name, p.Age, p.Score)
}

//设置分数
func (p *Graduate) SetScore(score int) {
	//业务判断
	p.Score = score
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中")
}

//Go中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性
func main() {
	var pupil = &Pupil{
		Name: "tom",
		Age:  10,
	}
	pupil.testing()
	pupil.SetScore(90)
	pupil.showInfo()
	//代码冗余了
	var graduate = &Graduate{
		Name: "marry",
		Age:  20,
	}
	graduate.testing()
	graduate.SetScore(88)
	graduate.showInfo()
}
