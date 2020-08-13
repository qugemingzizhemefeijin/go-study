package main

import (
	"fmt"
	"os"
)

//PathExists 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	//go判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断：
	//1) 如果返回的错误为nil，说明文件或文件夹存在
	//2) 如果返回的错误类型使用os.IsNotExists()判断为true，说明文件和文件夹不存在
	//3) 如果返回的错误为其它类型，则不确定是否存在

	b, err := PathExists("E:/1.txt")
	if b {
		fmt.Println("文件存在")
	} else if err == nil {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件不确定存不存在,err=", err)
	}
}
