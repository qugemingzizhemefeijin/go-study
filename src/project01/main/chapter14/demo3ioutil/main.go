package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//uoutil一次将整个文件读入到内存中，这种方式适合用于文件不大的情况
	//适用ioutil.ReadFile一次性将文件读取到位
	file := "E:/1.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	//把读取到的内容显示到终端
	fmt.Printf("%v\n", string(content)) //注意此处content是[]byte，需要转成string
	//我们没有显示的open文件，因此也不需要显示的Close文件，因为这些代码被封装在ReadFile函数内部了
}
