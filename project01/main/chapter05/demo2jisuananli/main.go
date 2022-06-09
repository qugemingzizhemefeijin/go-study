package main

import (
	"fmt"
	"math"
)

//如果成绩为100分时，奖励一辆BMW
//成绩为80，99时，奖励一台iphone7plus
//成绩为60-80时，奖励一个iPad;
//其它时，什么奖励也没有
//情从键盘输入期末成绩，并加以判断
func main() {
	fmt.Println("请输入成绩：")
	var cj int32
	_, _ = fmt.Scanln(&cj)

	if cj == 100 {
		fmt.Println("奖励一辆BMW")
	} else if cj >= 80 && cj < 100 {
		fmt.Println("奖励一台iPhone7Plus")
	} else if cj >= 60 && cj < 80 {
		fmt.Println("奖励一个iPad")
	} else {
		fmt.Println("都不及格，还想要毛啊")
	}

	//求ax2+bx+c=0方程的根。a,b,c分别为函数的参数，如果：b2-4ac>0。则有两个解：
	// b2-4ac=0，则有一个解，否则无解：
	// 提示1：x1=(-b+Sqrt(b2-4ac))/2a
	//        x2=(-b-Sqrt(b2-4ac))/2a
	// 提示2：math.Sqrt(num)；可以求平方根，需要引入math包
	var a, b, c float64 = 2.0, 4.0, 2.0
	m := b*b - 4*a*c

	if m > 0 {
		x1 := -b + math.Sqrt(m)/2*a
		x2 := -b - math.Sqrt(m)/2*a
		fmt.Printf("x1=%v,x2=%v\n", x1, x2)
	} else if m == 0 {
		x1 := -b + math.Sqrt(m)/2*a
		fmt.Printf("x1=%v\n", x1)
	} else {
		fmt.Println("无解")
	}

	//女方家长要嫁女儿，提出一定的条件：高：180cm以上；富：财富1000W以上；帅：是。条件从控制台输入。
	// 1) 如果三个条件都满足，则：我一定要嫁给他
	// 2) 如果三个条件有为真的情况，则：嫁吧，比上不足，比下有余
	// 3) 如果三个条件都不满足，则：不嫁
	var height int16
	var money int32
	var cool bool
	fmt.Println("请输入男方的条件：")
	fmt.Scanf("%d %d %t", &height, &money, &cool)
	fmt.Printf("男方身高%d,财富%d,帅气%t\n", height, money, cool)
	if height >= 180 && money >= 10000000 && cool {
		fmt.Println("我一定要嫁给他")
	} else if height >= 180 || money >= 10000000 || cool {
		fmt.Println("嫁吧，比上不足，比下有余")
	} else {
		fmt.Println("不嫁，滚")
	}
}
