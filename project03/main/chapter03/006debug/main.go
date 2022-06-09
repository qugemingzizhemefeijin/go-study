package main

import "fmt"

// 命令⾏进⼊包所在⽬ 录， 然后输⼊ dlv debug 命令进⼊调试
// dlv debug		目录执行此命令
// help				查看帮助
// break main.main	在main方法处设置断点
// breakpoints		查看所有断点
// vars main		查看全局变量
// continue			执行到下一个断点
// next				单步执行进入
// args				查看函数参数
// locals			查看局部变量
// stack			查看当前执行函数的栈帧信息
// goroutine		查看当前Goroutine信息
// goroutines		查看所有的Goroutine信息
// disassemble		查看函数对应的汇编代码
// regs				查看全部的寄存器状态

// 尴尬 Go version 1.14.7 is too old for this version of Delve (minimum supported version 1.16, suppress this error
// with --check-go-version=false)
func main() {
	nums := make([]int, 5)
	for i := 0; i < len(nums); i++ {
		nums[i] = i * i
	}
	fmt.Println(nums)
}
