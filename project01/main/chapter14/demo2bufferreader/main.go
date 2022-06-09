package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("E:/2.txt")
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	//档函数退出时，要即时的关闭file
	defer file.Close() //要及时关闭file，否则会有内存泄露

	//创建一个*Reader，是带缓冲的
	//默认缓冲大小为4096字节
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读取文件直到遇到\n字符才结束(本身换行符会被放到str中)
		if err != nil && err != io.EOF {    //如果读取文件失败
			fmt.Println("Read File err", err)
		}
		//输出内容
		fmt.Print(str)
		if err == io.EOF { //表示文件的末尾
			break
		}
	}
	fmt.Println("文件读取结束...")
}
