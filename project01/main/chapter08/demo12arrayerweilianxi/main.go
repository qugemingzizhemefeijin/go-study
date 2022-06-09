package main

import (
	"fmt"
)

//要求输入3个班级，每个班级5个同学，求出每个班的平均分以及所有班级平均分
func main() {
	//1.定义二维数组
	var scores [3][5]float64
	//2.循环的输入成绩
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	//fmt.Println(scores)
	//3.遍历输出成绩后的二维数组，统计平均分
	total := 0.0
	count := 0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
			count++
		}
		fmt.Printf("第%d班的总分为：%v，平均分：%v\n", i+1, sum, sum/float64(len(scores[i])))
		total += sum
	}
	fmt.Printf("所有班级的总分为：%.2f，平均分：%.2f\n", total, total/float64(count))
}
