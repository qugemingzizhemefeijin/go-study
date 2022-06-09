package main

/*
struct A {
	int i;
	float f;
	int type; // type 是 Go 语言的关键字，可以通过 _type来访问
	float _type; // 将屏蔽CGO对type成员的访问
	int size: 10; // 位字段无法访问
	float arr[]; // 零长的数组也无法访问
};
 */
import "C"
import "fmt"

// C语⾔的结构体、联合、枚举类型不能作为匿名成员被嵌⼊到Go语⾔的结构体中。在Go语⾔中，我们可以通过C.struct_xxx来访问C语⾔中定义的struct xxx结构体类型。
// 结构体的内存布局按照C语⾔的通⽤对⻬规则，在32位Go语⾔环境C语⾔结构体也按照32位对⻬规则，在64位Go语⾔环境按照64位的对⻬规则。对于指定了特殊对⻬规则的结构体，
// ⽆法在CGO中访问
func main() {
	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a.f)
	fmt.Println(a._type) // _type 对应 type
	//fmt.Println(a.size) // 错误：位字段无法访问
	//fmt.Println(a.arr) // 错误：零长的数组也无法访问
}
