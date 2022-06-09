package main

// 对于联合类型， 我们可以通过 C.union_xxx 来访问C语⾔中定义的 union xxx 类型。 但是Go语⾔中并不⽀持C语⾔联合类型， 它们会被转为对应⼤⼩的字节数组。

/*
#include <stdint.h>

union B1 {
	int i;
	float f;
};

union B2 {
	int8_t i8;
	int64_t i64;
};
 */
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var b1 C.union_B1;
	fmt.Printf("%T\n", b1)	// [4]uint8

	var b2 C.union_B2;
	fmt.Printf("%T\n", b2)	// [8]uint8

	// 如果需要操作C语⾔的联合类型变量， ⼀般有三种⽅法： 第⼀种是在C语⾔中定义辅助函数；
	// 第⼆种是通过Go语⾔的"encoding/binary"⼿⼯解码成员(需要注意⼤端⼩端问题)；
	// 第三种是使⽤ unsafe 包强制转型为对应类型(这是性能最好的⽅式)。

	fmt.Println("b1.i:", *(*C.int)(unsafe.Pointer(&b1)))
	fmt.Println("b1.f:", *(*C.float)(unsafe.Pointer(&b1)))

	// 虽然 unsafe 包访问最简单、 性能也最好， 但是对于有嵌套联合类型的情况处理会导致问题复杂化。
	// 对于复杂的联合类型， 推荐通过在C语⾔中定义辅助函数的⽅式处理。
}
