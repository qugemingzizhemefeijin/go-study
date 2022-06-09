package main

// 查看汇编代码
// go tool compile -S apkg.go

/*
go.cuinfo.packagename. SDWARFINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00                          ........
"".Id SNOPTRDATA size=8
        0x0000 37 25 00 00 00 00 00 00                          7%......
*/

// 其中 go tool compile 命令⽤于调⽤Go语⾔提供的底层命令⼯具， 其中 -S 参数表示输出汇编格式
// "".Id 对应Id变量符号， 变量的内存⼤⼩为8个字节。
// 变量的初始化内容为 37 25 00 00 00 00 00 00 ， 对应⼗六进制格式的0x2537， 对应⼗进制为9527。
// SNOPTRDATA是相关的标志， 其中NOPTR表示数据中不包含指针数据
var Id = 9527
