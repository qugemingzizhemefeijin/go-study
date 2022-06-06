package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//编写一个程序，将一个文件的内容，写入到另外一个文件。注：这两个文件已经存在了。
	//使用ioutil.ReadFile / iouitl.WriteFile来完成

	file1Path := "E:/1.txt"
	file2Path := "E:/2.txt"

	//1.首先读取文件内容
	content, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("read file err", err)
		return
	}

	fmt.Println(string(content))

	//2.写入到新的文件
	err = ioutil.WriteFile(file2Path, content, 0666)
	if err != nil {
		fmt.Println("write file err", err)
	}
}
