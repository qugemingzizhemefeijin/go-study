package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//copy 自己编写一个函数，接收两个文件路径
func copy(destFilePath, sourceFilePath string) bool {
	//获取reader
	srcFile, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println("open source file err", err)
		return false
	}
	//关闭
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	//获取writer
	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open dest file err", err)
		return false
	}
	//关闭
	defer destFile.Close()
	writer := bufio.NewWriter(destFile)

	//此处需要刷新缓冲区
	defer writer.Flush()

	//拷贝
	len, err := io.Copy(writer, reader)
	if err != nil && err != io.EOF {
		fmt.Println("copy file err", err)
	}
	fmt.Println("拷贝成功，共写入字节数:", len)
	return true
}

func main() {
	//将一张图片(任意文件)拷贝到另外一个目录下(如果目录不存在，需要创建目录)
	//copy函数是io包提供的
	b := copy("E:/xxx.jpg", "E:/113.jpg")
	if b {
		fmt.Println("拷贝成功")
	} else {
		fmt.Println("拷贝失败")
	}
}
