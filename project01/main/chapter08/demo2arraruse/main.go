package main

import "fmt"

func main() {
	//从终端循环输入5个成绩，保存到float64数组，并输出
	// var scores [5]float64
	// for i := 0; i < len(scores); i++ {
	// 	fmt.Printf("请输入第%d个元素的值\n", i+1)
	// 	fmt.Scanln(&scores[i])
	// }

	// fmt.Println("输出成绩=====")
	// for i := 0; i < len(scores); i++ {
	// 	fmt.Println(scores[i])
	// }

	//四种初始化数组的方式
	var numsArray01 [3]int = [3]int{1, 2, 3}
	var numsArray02 = [3]int{1, 2, 3}
	var numsArray03 = [...]int{6, 7, 8}
	//可以执行元素值对应的下标...
	var names = [3]string{1: "tom", 0: "jack", 2: "marry"}

	numsArray05 := [...]string{1: "chen", 0: "gang", 2: "wansui"}

	fmt.Println(numsArray01)
	fmt.Println(numsArray02)
	fmt.Println(numsArray03)
	fmt.Println(names)
	fmt.Println(numsArray05)

	//for range结构遍历
	for index, value := range numsArray02 {
		fmt.Printf("numsArray02[%d] = %v\n", index, value)
	}

	//遍历数组的时候，如果不想使用下标index，可以直接把下标index标为下划线
	for _, value := range numsArray03 {
		fmt.Printf("numsArray03 = %v\n", value)
	}
}
