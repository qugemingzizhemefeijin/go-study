package main

import (
	"fmt"
	"math/rand"
	"time"
)

//1.创建一个byte类型的26个元素的数组，分别方式'A-Z'。使用for循环访问所有元素并打印出来。提示：字符数据运算'A' + 1 = 'B'
func test01() {
	var letters [26]byte
	letters[0] = 'A'
	for i := 1; i < len(letters); i++ {
		letters[i] = letters[i-1] + 1
	}

	for _, value := range letters {
		fmt.Printf("%c ", value)
	}
	fmt.Println()
}

//2.请求出一个数组的最大值，并得到对应的小标
func test02() {
	numsArray := [...]int{3, 10, 5, 2, 7, 1}
	var idx, max int
	for index, value := range numsArray {
		if value > max {
			max = value
			idx = index
		}
	}
	fmt.Printf("数组numsArray=%v, max=%d, index=%d\n", numsArray, max, idx)
}

//3.请求出一个数组的和和平均值。
func test03() {
	numsArray := [...]int{10, 4, 9, 3, 6, 1, 8}
	var total int
	for _, value := range numsArray {
		total += value
	}
	fmt.Printf("数组元素和=%d, 平均值=%.2f\n", total, float64(total)/float64(len(numsArray)))
}

//4.随机生成5个数，并将其反转打印
func test04() {
	var numsArray [5]int
	for i := 0; i < len(numsArray); i++ {
		rand.Seed(time.Now().UnixNano() + int64(i)) //时间戳种子
		numsArray[i] = rand.Intn(100)
	}
	fmt.Printf("随机生成的%d个数字为:%v\n", len(numsArray), numsArray)

	//进行反转输出
	// for i := len(numsArray) - 1; i >= 0; i-- {
	// 	fmt.Printf("%d ", numsArray[i])
	// }
	// fmt.Println()

	//直接可以交换
	temp := 0
	for i := 0; i < len(numsArray)/2; i++ {
		temp = numsArray[len(numsArray)-1-i]
		numsArray[len(numsArray)-1-i] = numsArray[i]
		numsArray[i] = temp
	}
	fmt.Printf("交换之后的数组:%v\n", numsArray)
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
