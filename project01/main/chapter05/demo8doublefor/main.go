package main

import (
	"fmt"
)

func main() {
	//1.统计3个班成绩情况，每个班有5名同学，求出各个班的平均分和所有班级的平均分
	//2.统计三个班及格人数，每个班5名同学
	// var score, count float64
	// var classNum int = 3
	// var stuNum int = 5
	// var passCount int = 0
	// for j := 1; j <= classNum; j++ {
	// 	sum := 0.0
	// 	for i := 1; i <= stuNum; i++ {
	// 		var s float64
	// 		fmt.Printf("请你输入第%d班级第%d个学生的成绩 \n", j, i)
	// 		fmt.Scanln(&s)

	// 		score += s
	// 		sum += s
	// 		count++

	// 		if s >= 60 {
	// 			passCount++
	// 		}
	// 	}
	// 	fmt.Printf("班级%d平均分:%v\n", j, sum/float64(stuNum))
	// }
	// fmt.Printf("所有班级的平均分:%v\n", score/count)
	// fmt.Printf("及格人数为：%v \n", passCount)
	//3.打印金字塔经典案例
	// var i int = 5
	// for n := 1; n <= i; n++ {
	// 	for m := 0; m < n; m++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	//*数量 = 2 * 层数 - 1
	// 空格数量 = 总层数 - 当前层数
	var i int = 5
	for n := 1; n <= i; n++ {
		for k := 1; k <= i-n; k++ {
			fmt.Print(" ")
		}
		for m := 1; m <= n*2-1; m++ {
			if m == 1 || m == n*2-1 || n == i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
