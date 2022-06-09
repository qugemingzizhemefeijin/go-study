package main

// 导出⾮main包的函数
// C函数必须是在main包导出， 然后才能在⽣成的头⽂件包含声明的语句。
// 但是很多时候我们可能更希望将不同类型的导出函数组织到不同的Go包中， 然后统⼀导出为⼀个静态库或动态库。

// 要实现从是从⾮main包导出C函数，或者是多个包导出C函数（因为只能有⼀个main包），我们需要⾃⼰提供导出C函数对应的头⽂件
// （因为CGO⽆法为⾮main包的导出函数⽣成头⽂件） 。

import "C"

import (
	_ "./numbers"
	"fmt"
)

// 其中我们导⼊了number⼦包， 在number⼦包中有导出的C函数number_add_mod， 同时我们在main包也导出了goPrintln函数。
// 通过以下命令创建C静态库：
// go build -buildmode=c-archive -o main.a

// 这时候在⽣成main.a静态库的同时， 也会⽣成⼀个main.h头⽂件。 但是main.h头⽂件中只有main包中导出的goPrintln函数的声明，
// 并没有number⼦包导出函数的声明。 其实number_add_mod函数在⽣成的C静态库中是存在的， 我们可以直接使⽤。

// gcc -o a.out _test_main.c main.a
// ./a.out

// 我们并没有包含CGO⾃动⽣成的main.h头⽂件，⽽是通过⼿⼯⽅式声明了goPrintln和number_add_mod两个导出函数。这样我们就实现了从多个Go包导出C函数了。
func main() {
	println("Done")
}

//export goPrintln
func goPrintln(s *C.char) {
	fmt.Println("goPrintln:", C.GoString(s))
}
