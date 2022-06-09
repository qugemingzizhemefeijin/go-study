package main

import (
	"flag"
	"fmt"
)

func main() {
	//编写一段代码，获取命令行各个参数。需要拿到u port等等
	var user, pwd, host string
	var port int

	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "host", "localhost", "主机地址,默认为空")
	flag.IntVar(&port, "port", 3306, "端口号,默认为3306")

	//这里有一个非常重要的操作，转换，必须调用该方法。这个方法调用时间点必须在FlagSet的所有标签都定义之后，程序访问这些标签之前。
	flag.Parse()

	fmt.Printf("用户名=%v,密码=%v,主机名=%v,端口=%d\n", user, pwd, host, port)
}
