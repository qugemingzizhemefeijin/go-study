package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//Node ...
type Node struct {
	row, col, val int
}

//稀疏数组
//当一个数组中大部分元素为0，或者为同一个值的数组时，可以使用稀疏数组来保存该数组
//稀疏数组的处理方法是：
//1) 记录数组一共有几行几列，有多少个不同的值
//2) 把具有不同值的元素的行列及值记录在一个小规模的数组中，从而缩小程序的规模

//应用实例
//1) 使用稀疏数组，来保留类似上面提到的二维数组（棋盘、地图等）
//2) 把稀疏数组存盘，并且可以从新恢复原来的二维数组数
func main() {
	//1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子

	//2.输出看看原始的数组
	printChessMap(chessMap)

	//3. 转成稀疏数组
	// 思路
	//(1). 遍历 chessMap ，如果我们发现有一个元素的值，不等于0，则创建一个Node结构体
	//(2). 将其放入到切片中
	var sparseArr []Node

	//标准的一个稀疏数组应该还含有表示记录原始的二维数组的规模（行和列，默认值）
	valNode := Node{
		row: len(chessMap),
		col: len(chessMap[0]),
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, n := range v {
			if n != 0 {
				//创建一个节点
				valNode = Node{
					row: i,
					col: j,
					val: n,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	fmt.Println("当前的稀疏数组是...")
	//输出稀疏数组
	for i, node := range sparseArr {
		fmt.Printf("%d: %d\t%d\t%d\n", i, node.row, node.col, node.val)
	}

	//将这个稀疏数组存盘 E:/chessmap.data
	//fmt.Printf("sparseArr address pointer %p \n", sparseArr)
	writeFile(sparseArr)

	//如何恢复原始的数组

	//1. 打开这个存盘的文件，通过文件恢复原始数据 E:/chessmap.data

	//2.这里使用稀疏数组恢复
	sparseArr2 := readFile()

	//先创建一个原始数组
	var chessMap2 [11][11]int
	//遍历 sparseArr
	for i, node := range sparseArr2 {
		if i == 0 {
			continue
		}
		chessMap2[node.row][node.col] = node.val
	}

	fmt.Println("恢复后的原始数据")
	//fmt.Printf("chessMap address %p \n", &chessMap2)
	printChessMap(chessMap2)
}

//将文件数组读取到稀疏数组中
func readFile() (sparseArr []Node) {
	file, err := os.Open("E:/chessmap.data")
	if err != nil {
		fmt.Println("Open error = ", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("读取异常了,error=", err)
			break
		}
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			break
		}

		//此处初步恢复数据
		s := strings.Split(str, " ")
		row, _ := strconv.ParseInt(s[0], 10, 32)
		col, _ := strconv.ParseInt(s[1], 10, 32)
		val, _ := strconv.ParseInt(s[2], 10, 32)
		node := Node{
			row: int(row),
			col: int(col),
			val: int(val),
		}
		sparseArr = append(sparseArr, node)
	}
	return
}

//将稀疏数组写入到文件中
func writeFile(sparseArr []Node) {
	//fmt.Printf("writeFile sparseArr address pointer %p \n", sparseArr)

	file, err := os.OpenFile("E:/chessmap.data", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("OpenFile error = ", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, node := range sparseArr {
		writer.WriteString(fmt.Sprintf("%d %d %d\n", node.row, node.col, node.val))
	}
	writer.Flush()
}

func printChessMap(chessMap [11][11]int) {
	//fmt.Printf("chessMap address %p \n", &chessMap)
	for _, v := range chessMap {
		for _, n := range v {
			fmt.Printf("%d\t", n)
		}
		fmt.Println()
	}
}
