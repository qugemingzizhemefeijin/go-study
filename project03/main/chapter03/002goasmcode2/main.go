package main

import pkg "./pkg"

// pkg.go
/*
go.cuinfo.packagename. SDWARFINFO dupok size=0
        0x0000 70 6b 67                                         pkg
go.string."gopher" SRODATA dupok size=6
        0x0000 67 6f 70 68 65 72                                gopher
"".Name SDATA size=16
        0x0000 00 00 00 00 00 00 00 00 06 00 00 00 00 00 00 00  ................
        rel 0+8 t=1 go.string."gopher"+0
*/

// 输出中出现了⼀个新的符号go.string."gopher"， 根据其⻓度和内容分析可以猜测是对应底层的"gopher"字符串数据。
// 因为Go语⾔的字符串并不是值类型， Go字符串其实是⼀种只读的引⽤类型。
// 如果多个代码中出现了相同的"gopher"只读字符串时， 程序链接后可以引⽤的同⼀个符号go.string."gopher"。
// 因此， 该符号有⼀个SRODATA标志表示这个数据在只读内存段， dupok表示出现多个相同标识符的数据时只保留⼀个就可以了。
// ⽽真正的Go字符串变量Name对应的⼤⼩却只有16个字节了。
// 其实Name变量并没有直接对应“gopher”字符串， ⽽是对应16字节⼤⼩的reflect.StringHeader结构体
/*
type reflect.StringHeader struct {
	Data uintptr
	Len  int
}
 */
// 从汇编角度看，Name 变量其实对应的是 reflect.StringHeader 结构体类型。前 8 个字节对应底层真实字符串数据的指针，
// 也就是符号 go.string."gopher" 对应的地址。后 8 个字节对应底层真实字符串数据的有效长度，这里是 6 个字节。

// 因为在 Go 汇编语言中，go.string."gopher" 不是一个合法的符号，
// 因此我们无法通过手工创建（这是给编译器保留的部分特权，因为手工创建类似符号可能打破编译器输出代码的某些规则）。
// 因此我们新创建了一个 ·NameData 符号表示底层的字符串数据。然后定义 ·Name 符号内存大小为 16 字节，其中前 8 个字节用
// ·NameData 符号对应的地址初始化，后 8 个字节为常量 6 表示字符串长度。

// 第一次运行： pkgpath.NameData: missing Go type information for global symbol: size 8 错误提示汇编中定义的 NameData 符号没有类型信息。

// 其实 Go 汇编语言中定义的数据并没有所谓的类型，每个符号只不过是对应一块内存而已，因此 NameData 符号也是没有类型的。
// 但是 Go 语言是带垃圾回收器的语言，Go 汇编语言工作在这个自动垃圾回收体系框架内。当 Go 语言的垃圾回收器在扫描到 NameData 变量的时候，
// 无法知晓该变量内部是否包含指针，因此就出现了这种错误。错误的根本原因并不是 NameData 没有类型，而是 NameData 变量没有标注是否会含有指针信息。

// 通过给 NameData 变量增加一个 NOPTR 标志，表示其中不会包含指针数据可以修复该错误：
// #include "textflag.h"
// GLOBL ·NameData(SB),NOPTR,$8

// 通过给 ·NameData 增加 NOPTR 标志的方式表示其中不含指针数据。
// 我们也可以通过给 ·NameData 变量在 Go 语言中增加一个不含指针并且大小为 8 个字节的类型来修改该错误：

// 我们将 NameData 声明为长度为 8 的字节数组。编译器可以通过类型分析出该变量不会包含指针，因此汇编代码中可以省略 NOPTR 标志。
// 现在垃圾回收器在遇到该变量的时候就会停止内部数据的扫描。

// 在这个实现中，Name 字符串底层其实引用的是 NameData 内存对应的 “gopher” 字符串数据。因此，如果 NameData 发生变化，Name 字符串的数据也会跟着变化。

// 当然这和字符串的只读定义是冲突的，正常的代码需要避免出现这种情况。最好的方法是不要导出内部的 NameData 变量，这样可以避免内部数据被无意破坏。

// 在用汇编定义字符串时我们可以换一种思维：将底层的字符串数据和字符串头结构体定义在一起，这样可以避免引入 NameData 符号：
/*
GLOBL ·Name(SB),$24

DATA ·Name+0(SB)/8,$·Name+16(SB)
DATA ·Name+8(SB)/8,$6
DATA ·Name+16(SB)/8,$"gopher"
 */

// 在新的结构中，Name 符号对应的内存从 16 字节变为 24 字节，多出的 8 个字节存放底层的 “gopher” 字符串。
// ·Name 符号前 16 个字节依然对应 reflect.StringHeader 结构体：
// Data 部分对应 $·Name+16(SB)，表示数据的地址为 Name 符号往后偏移 16 个字节的位置；
// Len 部分依然对应 6 个字节的长度。
// 这是 C 语言程序员经常使用的技巧。

// 定义字符串变量
func main() {
	println(pkg.Name)

	pkg.NameData[0] = '?'
	println(pkg.Name)
}
