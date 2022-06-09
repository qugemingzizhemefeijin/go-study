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

