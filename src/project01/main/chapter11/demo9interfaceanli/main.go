package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//实现对Hero结构体切片的排序：sort.Sort(data Interface)

//Hero 先声明结构体
type Hero struct {
	Name string
	Age  int
}

//2.声明一个Hero结构体切片类型
type HeroSlice []Hero

//3.实现Sort的接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

//这个决定你使用什么标准进行排序
//1.按Hero的年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	//与上面等价
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	//定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	//要求对intSlice进行排序
	//1.冒泡排序
	//2.也可以使用系统系统的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//请大家编写对一个结构体切片进行排序
	//1.仍然可以进行冒泡排序
	//2.也可以使用系统提供的方法

	// var heroes = HeroSlice{
	// 	Hero{"萨姆", 200},
	// 	Hero{"牛魔王", 250},
	// 	Hero{"孙悟空", 100},
	// 	Hero{"芭蕉精", 50},
	// }
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	sort.Sort(heroes)

	fmt.Println(heroes)
}
