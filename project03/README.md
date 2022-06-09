### GO高级编程 - 编程源码

最初CGO是为了达到⽅便从Go语⾔函数调⽤C语⾔函数（⽤C语⾔实现Go语⾔声明的函数）以复⽤C语⾔资源这⼀⽬的⽽出现的（因为C语⾔还会涉及回调函数，⾃然也会涉及到从C语⾔函数调⽤Go语⾔函数（⽤Go语⾔实现C语⾔声明的函数））。现在，它已经演变为C语⾔和Go语⾔双向通讯的桥梁。要想利⽤好CGO特性，⾃然需要了解此⼆语⾔类型之间的转换规则

| C 语言类型               | CGO 类型      | Go 语言类型 |
| --- | --- | --- |
| char                   | C.char      | byte |
| singed char            | C.schar     | int8 |
| unsigned char          | C.uchar     | uint8 |
| short                  | C.short     | int16 |
| unsigned short         | C.ushort     | uint16 |
| int                    | C.int       | int32 |
| unsigned int           | C.uint      | uint32 |
| long                   | C.long      | int32 |
| unsigned long          | C.ulong     | uint32 |
| long long int          | C.longlong  | int64 |
| unsigned long long int | C.ulonglong | uint64 |
| float                  | C.float     | float32 |
| double                 | C.double    | float64 |
| size_t                 | C.size_t    | uint |

除了 `GoInt` 和 `GoUint` 之外，我们并不推荐直接访问 `GoInt32`、`GoInt64` 等类型。更好的做法是通过 C 语言的 C99 标准引入的 `<stdint.h>` 头文件。为了提高 C 语言的可移植性，在 `<stdint.h>` 文件中，不但每个数值类型都提供了明确内存大小，而且和 Go 语言的类型命名更加一致。

| C 语言类型 | CGO 类型     | Go 语言类型 |
| --- | --- | --- |
| int8_t   | C.int8_t   | int8 |
| uint8_t  | C.uint8_t  | uint8 |
| int16_t  | C.int16_t  | int16 |
| uint16_t | C.uint16_t | uint16 |
| int32_t  | C.int32_t  | int32 |
| uint32_t | C.uint32_t | uint32 |
| int64_t  | C.int64_t  | int64 |
| uint64_t | C.uint64_t | uint64 |

### 编译参数

- `CFLAGS`对应`C`语言编译参数(以 .c 后缀名)
- `CPPFLAGS`对应`C/C++`代码编译参数(.c,.cc,.cpp,.cxx)
- `CXXFLAGS`对应纯`C++`编译参数(.cc,.cpp,*.cxx)
- `LDFLAGS`链接参数主要包含要链接库的检索目录和要链接库的名字。链接库不支持相对路径，我们必须为链接库指定绝对路径。`cgo`中的`${SRCDIR}`为当前目录的绝对路径。

### dlv debug调式

| 命令 | 描述 |
| --- | --- |
| dlv debug	| 目录执行此命令 |
| help | 查看帮助 |
| break main.main | 在main方法处设置断点 |
| breakpoints | 查看所有断点 |
| vars main | 查看全局变量 |
| continue | 执行到下一个断点 |
| next | 单步执行进入 |
| args | 查看函数参数 |
| locals | 查看局部变量 |
| stack | 查看当前执行函数的栈帧信息 |
| goroutine | 查看当前Goroutine信息 |
| goroutines | 查看所有的Goroutine信息 |
| disassemble | 查看函数对应的汇编代码 |
| regs | 查看全部的寄存器状态 |
