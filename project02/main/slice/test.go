package main

import "fmt"

//Go的slice类型中包含了一个array指针以及len和cap两个int类型的成员。
//Go中的参数传递实际都是值传递，将slice作为参数传递时，函数中会创建一个slice参数的副本，这个副本同样也包含array,len,cap这三个成员。
//副本中的array指针与原slice指向同一个地址，所以当修改副本slice的元素时，原slice的元素值也会被修改。但是如果修改的是副本slice的len和cap时，原slice的len和cap仍保持不变。
//如果在操作副本时由于扩容操作导致重新分配了副本slice的array内存地址，那么之后对副本slice的操作则完全无法影响到原slice，包括slice中的元素。

//https://www.bilibili.com/read/cv6468706/

//SliceRise ...
func SliceRise(s []int) {
	fmt.Printf("s=%v,len=%d,cap=%d,addr=%p,s[0]=%p\n", s, len(s), cap(s), s, &s[0])
	s = append(s, 0)
	fmt.Printf("s=%v,len=%d,cap=%d,addr=%p,s[0]=%p\n", s, len(s), cap(s), s, &s[0])
	for i := range s {
		fmt.Printf("s[%d]=%d\n", i, s[i])
		s[i]++
	}
}

//SlicePrint ...
func SlicePrint() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	SliceRise(s1)
	fmt.Println("===================================================")
	fmt.Printf("s2=%v,len=%d,cap=%d,addr=%p,s[0]=%p\n", s2, len(s2), cap(s2), s2, &s2[0])
	SliceRise(s2)
	fmt.Printf("s2=%v,len=%d,cap=%d,addr=%p,s[0]=%p\n", s2, len(s2), cap(s2), s2, &s2[0])
	fmt.Println(s1, s2)
}

func main() {
	SlicePrint()
}
