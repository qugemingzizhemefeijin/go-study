package main

import "fmt"

/*
切片就是一种简化版的动态数组。因为动态数组的长度是不固定，切片的长度自然也就不能是类型的组成部分了。
切片的底层结构定义
type SliceHeader struct {
	Data uintptr
	Len int
	Cap int
}
*/
func main() {
	var (
		a []int               // nil切片，和nil相等，一般用来表示一个不存在的切片
		b = []int{}           // 空切片，喝nil不相等，一般用来表示一个空的集合
		c = []int{1, 2, 3}    // 有3个元素的切片，len和cap都为3
		d = c[:2]             // 有2个元素的切片，len为2，cap为3
		e = c[0:2:cap(c)]     // 有2个元素的切片，len为2，cap为3 元素内容：1,2
		f = c[:0]             // 有0个元素的切片，len为0，cap为3
		g = make([]int, 3)    // 有3个元素的切片，len和cap都为3
		h = make([]int, 2, 3) // 有2个元素的切片，len为2，cap为3
		j = make([]int, 0, 3) // 有0个元素的切片，len为0，cap为3
	)

	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}

	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}

	fmt.Println("=========================")

	// 容量不足的情况下，append的操作会导致重新分配内存，可能导致巨大的内存分配和复制数据代价。
	// 即使容量足够，依然需要用append函数的返回值来更新切片本身，因为新切片的长度已经发生变化。

	d = append(d, 1)                 // 追加1个元素
	d = append(d, 1, 2, 3)           // 追加多个元素，手写解包方式
	d = append(d, []int{1, 2, 3}...) // 追加一个切片，切片需要解包

	for i, v := range d {
		fmt.Printf("d[%d]: %d\n", i, v)
	}

	fmt.Println("=========================")

	// 除了在切片的尾部追加，还可以在切片的开头添加元素
	// 在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。因为从切片的开头添加元素的性能要比尾部追加元素的性能差很多。
	d = append([]int{0}, d...)          // 在开头添加1个元素
	d = append([]int{-3, -2, -1}, d...) // 在开头添加1个切片

	for i, v := range d {
		fmt.Printf("d[%d]: %d\n", i, v)
	}

	fmt.Println("=========================")

	// 由于append函数返回新的切片，也就是它支持链式操作。
	e = append(e[:1], append([]int{3}, e[1:]...)...)       // 在第2个位置插入3
	e = append(e[:2], append([]int{1, 2, 3}, e[2:]...)...) // 在第2个位置插入切片

	for i, v := range e {
		fmt.Printf("e[%d]: %d\n", i, v)
	}

	// 每个添加操作的第二个append调用都会创建一个临时切片，并将内容复制到新创建的切片中，然后将临时创建的切片再追加到e[1:]。
	// 可以使用copy和append组合来避免创建中间的临时切片

	fmt.Println("=========================")

	f = append(f, 1, 2, 3) // 初始化
	f = append(f, 0)       // 扩展1个空间
	copy(f[3:], f[2:])     // 将f[2]向后移动1个位置
	f[2] = 4
	for i, v := range f {
		fmt.Printf("f[%d]: %d\n", i, v)
	}

	fmt.Println("=========================")

	x := []int{7, 8, 9}
	g = append(g, x...)       // 为x切片扩展足够的空间
	copy(g[2+len(x):], g[2:]) // g[2:]向后移动len(x)个位置
	copy(g[2:], x)            // 复制新添加的切片

	for i, v := range g {
		fmt.Printf("g[%d]: %d\n", i, v)
	}

	fmt.Println("=========================")

	// 删除切片元素根据位置有三种情况：从开头位置删除，从中间位置删除，从尾部山粗。其中删除切片的尾部元素最快

	h = []int{1, 2, 3, 4}
	h = h[:len(h)-1] // 从尾部删除1个元素
	h = h[:len(h)-2] // 从尾部删除2个元素

	for i, v := range h {
		fmt.Printf("h[%d]: %d\n", i, v)
	}

	// 删除开头的元素也可以直接移动数据指针
	h = []int{1, 2, 3, 4}
	h = h[1:] // 删除开头1个元素
	h = h[2:] // 删除开头2个元素

	for i, v := range h {
		fmt.Printf("h[%d]: %d\n", i, v)
	}

	// 删除开头的元素也可以不移动数据指针，可以将数据向开头移动，可以用append原地完成
	// 所以原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化
	h = []int{1, 2, 3, 4}
	h = append(h[:0], h[1:]...) // 删除开头1个元素
	h = append(h[:0], h[2:]...) // 删除开头2个元素

	for i, v := range h {
		fmt.Printf("h[%d]: %d\n", i, v)
	}

	// 也可以用copy完成删除开头的元素
	h = []int{1, 2, 3, 4}
	h = h[:copy(h, h[1:])] // 删除开头1个元素
	h = h[:copy(h, h[2:])] // 删除开头2个元素

	for i, v := range h {
		fmt.Printf("h[%d]: %d\n", i, v)
	}

	// 对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用append或copy原地完成
	j = []int{1, 2, 3, 4, 5}
	j = append(j[:1], j[2:]...) // 删除中间1个元素
	j = append(j[:1], j[3:]...) // 删除中间2个元素

	for i, v := range j {
		fmt.Printf("j[%d]: %d\n", i, v)
	}

	j = []int{1, 2, 3, 4, 5}
	j = j[:1+copy(j[1:], j[2:])] // 删除中间1个元素
	j = j[:1+copy(j[1:], j[3:])] // 删除中间2个元素

	for i, v := range j {
		fmt.Printf("j[%d]: %d\n", i, v)
	}

	// 避免切片内存泄露
	// 返回的时候，如果只是需要切片的一部分，可以将那部分拷贝到新的切片中返回。return append([]byte{}, b...)
	// 删除的时候，因为直接删除，其实底层数组还是继续引用者，最好可以将删除的索引指向nil，当然如果切片存在的周期很短，不需要刻意来做这个。

}
