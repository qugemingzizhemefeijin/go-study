package main

import (
	"fmt"
	"os"
)

func main() {
	//打开一个文件
	file, err := os.Open("E:/2.txt")
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}

	//输出下文件，看看文件是什么
	fmt.Printf("file=%v\n", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	} else {
		fmt.Println("close file success")
	}
}
