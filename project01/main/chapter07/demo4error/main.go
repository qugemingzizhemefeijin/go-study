package main

import (
	"errors"
	"fmt"
)

func test() {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		//recover()是内置函数，可以捕获到异常
		if err := recover(); err != nil { //说明捕获到错误
			fmt.Println("err", err)
		}
	}()
	//测试如果defer中发生错误，是否会影响上面的defer(测试了之后，好像go保证每个defer都能被执行)
	defer func() {
		fmt.Println("+++++++++++++++++")
		var i, j int = 10, 0
		fmt.Println("====", i/j)
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//函数去读取配置文件init.conf的系你想
//如果文件名传入不正确，我们就返回一个自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	}
	//返回一个自定义错误
	return errors.New("读取文件错误")
}

func test02() {
	err := readConf("config1.ini")
	if err != nil {
		//如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行")
}

//Golang的错误处理机制
//Go语言追求简洁优雅，所以，Go语言不支持传统的trt...catch...finally这种处理
//Go中引入的处理方式为：defer，panic，recover
//这几个异常的使用场景可以这么简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理
//Go中也支持自定义错误，使用errors.New和panic内置喊出。
//1.errors.New("错误说明")，会返回一个error类型的值，表示一个错误
//2.panic内置函数，接收一个interface{}类型的值(也就是任意值)作为参数，可以接收error类型的变量，输出错误信息，并退出程序。
func main() {
	// test()
	// fmt.Println("下面的代码和逻辑...")
	// for {
	// 	fmt.Println("main()下面的代码")
	// 	time.Sleep(time.Second)
	// }

	//测试自定义错误的使用
	test02()
	fmt.Println("main()下面的代码")
}
