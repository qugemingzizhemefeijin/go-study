package main

// 先创建 main.go 文件，创建并初始化字符串变量，同时声明 main 函数：

var helloworld = "你好, 世界"

/*
TEXT ·main(SB), $16-0
	MOVQ ·helloworld+0(SB), AX; MOVQ AX, 0(SP)
	MOVQ ·helloworld+8(SB), BX; MOVQ BX, 8(SP)
	CALL runtime·printstring(SB)
	CALL runtime·printnl(SB)
	RET

TEXT ·main(SB), $16-0 用于定义 main 函数
其中 $16-0 表示 main 函数的帧大小是 16 个字节（对应 string 头部结构体的大小，用于给 runtime·printstring 函数传递参数），
0 表示 main 函数没有参数和返回值。

main 函数内部通过调用运行时内部的 runtime·printstring(SB) 函数来打印字符串。
然后调用 runtime·printnl 打印换行符号。

Go 语言函数在函数调用时，完全通过栈传递调用参数和返回值。
先通过 MOVQ 指令，将 helloworld 对应的字符串头部结构体的 16 个字节复制到栈指针 SP 对应的 16 字节的空间，然后通过 CALL 指令调用对应函数。
最后使用 RET 指令表示当前函数返回。
 */

// missing function body 跑不通这个例子...
func main()
