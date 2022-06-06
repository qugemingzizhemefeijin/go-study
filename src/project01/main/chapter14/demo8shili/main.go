package main

import (
	"fmt"
	"io/ioutil"
)

//CharCount 定义一个结构体，用户保存统计的结构
type CharCount struct {
	ChCount    int //记录英文个数
	NumCount   int //记录数字个数
	SpaceCount int //记录空格个数
	OtherCount int //记录其它字符个数
}

//统计英文、数字、空格和其他字符数量
func main() {
	content, err := ioutil.ReadFile("E:/1.txt")
	if err != nil {
		fmt.Println("Read File Error = ", err)
		return
	}

	var count CharCount
	str := string(content)
	for _, v := range str {
		switch {//switch也可以不需要type类型
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			count.ChCount++
		case v >= '0' && v <= '9':
			count.NumCount++
		case v == ' ' || v == '\t':
			count.SpaceCount++
		default:
			count.OtherCount++
		}
	}

	fmt.Printf("英文=%v,数字=%v,空格=%v,其它=%v\n", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
