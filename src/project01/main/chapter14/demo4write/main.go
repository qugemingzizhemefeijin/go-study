package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//writerAndCreateFile 1) 创建一个新文件，写入内容5句"hello,gardon"
func writerAndCreateFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open File err", err)
		return
	}
	//关闭file
	defer func() {
		file.Close()
		fmt.Println("关闭file")
	}()

	//准备写入
	str := "hello,gardon\n"
	//使用时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为Writer是带缓存的，在调用WriteString方法时，其实内容是先写入到缓存的，所以需要调用Flush()方法刷新到磁盘
	writer.Flush()
}

//writerCoverFile 2) 打开一个存在的文件，将原来的内容覆盖成新的内容10句 "你好，小明"
func writerCoverFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Open File err", err)
		return
	}
	//关闭file
	defer file.Close()

	//准备写入
	str := "你好，小明\n"
	//使用时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	//因为Writer是带缓存的，在调用WriteString方法时，其实内容是先写入到缓存的，所以需要调用Flush()方法刷新到磁盘
	writer.Flush()
}

//writerAppendFile 3) 打开一个存在的文件，在原来的内容追加内容"ABCI ENGLISH!"
func writerAppendFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open File err", err)
		return
	}
	//关闭file
	defer file.Close()

	//使用时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	writer.WriteString("ABCI ENGLISH!\n")

	//因为Writer是带缓存的，在调用WriteString方法时，其实内容是先写入到缓存的，所以需要调用Flush()方法刷新到磁盘
	writer.Flush()
}

//writerAndReaderFile 4) 打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello,北京!"
func writerAndReaderFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open File err", err)
		return
	}
	//关闭file
	defer file.Close()

	fmt.Println("===================================")
	//先读取
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("read file err", err)
			break
		}
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}

	//使用时，使用带缓存的*Writer
	str := "hello,北京!\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为Writer是带缓存的，在调用WriteString方法时，其实内容是先写入到缓存的，所以需要调用Flush()方法刷新到磁盘
	writer.Flush()
}

//readFile 读取文件内容
func readFile(filePath string) {
	fmt.Println("===================================")
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("文件读取失败!")
		return
	}
	fmt.Println(string(content))
}

func main() {
	//OpenFile(name string, flag int, perm FileMode)
	// name  文件名
	// flag  文件打开模式
	//O_RDONLY int = syscall.O_RDONLY // 只读模式打开
	//O_WRONLY int = syscall.O_WRONLY // 只写模式打开
	//O_RDWR   int = syscall.O_RDWR   // 读写模式打开
	//O_APPEND int = syscall.O_APPEND // 追加模式
	//O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	//O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	//O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步IO
	//O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	// perm  文件默认权限
	//ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
	//ModeAppend                                     // a: 只能写入，且只能写入到末尾
	//ModeExclusive                                  // l: 用于执行
	//ModeTemporary                                  // T: 临时文件(非备份文件)
	//ModeSymlink                                    // L: 符号链接(不是快捷方式文件)
	//ModeDevice                                     // D: 设备
	//ModeNamedPipe                                  // p: 命名管道(FIFO)
	//ModeSocket                                     // S: Unix域socket
	//ModeSetuid                                     // u: 表示文件具有其创建者用户ID权限
	//ModeSetgid                                     // g: 表示文件具有其创建者组ID的权限
	//ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
	//ModeSticky                                     // t: 只有root/创建者能删除/移动文件
	// 覆盖所有类型位(用于通过&获取类型位)，对普通文件，所有这些位都不应该背设置
	//ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
	//ModePerm FileMode = 0777 // 覆盖所有Unix权限位(用于通过&获取类型位)

	filePath := "E:/1.txt"
	//1) 创建一个新文件，写入内容5句"hello,gardon"
	writerAndCreateFile(filePath)
	readFile(filePath)
	//2) 打开一个存在的文件，将原来的内容覆盖成新的内容10句 "你好，小明"
	writerCoverFile(filePath)
	readFile(filePath)
	//3) 打开一个存在的文件，在原来的内容追加内容"ABCI ENGLISH!"
	writerAppendFile(filePath)
	readFile(filePath)
	//4) 打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello,北京!"
	writerAndReaderFile(filePath)
	readFile(filePath)
	//使用os.OpenFile(),bufio.NewWriter() *Writer的方法WriteString完成上面的任务

}
