package main

/*
enum C {
	ONE,
	TWO,
};
 */
import "C"
import "fmt"

// 枚举类型， 我们可以通过 C.enum_xxx 来访问C语⾔中定义的 enum xxx 结构体类型。
// 在C语⾔中， 枚举类型底层对应 int 类型， ⽀持负数类型的值。 我们可以通过 C.ONE 、 C.TWO 等直接访问定义的枚举值。
func main() {
	var c C.enum_C = C.TWO
	fmt.Println(c)
	fmt.Println(C.ONE)
	fmt.Println(C.TWO)
}
