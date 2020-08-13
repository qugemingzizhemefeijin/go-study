package main

import (
	"fmt"
	"sort"
)

//测试是否是引用
func modify(map1 map[int]int) {
	map1[10] = 999
}

//Student 定义一个学生结构体
type Student struct {
	Name    string
	Age     int
	Address string
}

//map的key支持bool、数字、string、指针、channel，还可以是只包含前面几个类型的接口、结构体、数组
//通常来说valuetype的类型为数字、string、map、struct等；
//slice、map还有function不可以做为key，因为这几个没法用==来判断
func main() {
	//声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用
	//注意：这一点跟数组是不一样的
	//第一种，先声明，再分配空间
	var a map[string]string
	//var a map[string]map[string]string
	//在使用map前，需要先make，make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江" //这个必须要写在make之后
	a["no2"] = "吴用"
	a["no3"] = "吴用"
	fmt.Println(a)
	//第二种，声明即赋值
	var b = make(map[string]string)
	b["no1"] = "北京"
	b["no2"] = "上海"
	fmt.Println(b)

	city := b["no1"]
	fmt.Println(city)

	//第三种，直接赋值  前面这个map[string]string可以省掉的
	var c map[string]string = map[string]string{"no1": "关胜", "no2": "卢俊义"}
	fmt.Println(c)

	//第四种
	d := map[string]string{"no1": "小明", "no2": "朱武"}
	fmt.Println(d)

	//我们要存放3个学生信息，每个学生有name和sex信息
	var students = make(map[string]map[string]string)
	students["stu01"] = make(map[string]string, 3)
	students["stu01"]["name"] = "tom"
	students["stu01"]["ses"] = "男"
	students["stu01"]["address"] = "北京朝阳区"

	students["stu02"] = make(map[string]string, 3)
	students["stu02"]["name"] = "marry"
	students["stu02"]["ses"] = "女"
	students["stu02"]["address"] = "上海静安区"

	students["stu03"] = make(map[string]string, 3)
	students["stu03"]["name"] = "jack"
	students["stu03"]["ses"] = "男"
	students["stu03"]["address"] = "英国伦敦"

	//其实也可以如下来初始化，感觉会更方便能看一点
	//students["no1"] = map[string]string{"name": "小明", "sex": "男"}
	//students["no2"] = map[string]string{"name": "小红", "sex": "女"}
	//students["no3"] = map[string]string{"name": "小花", "sex": "女"}
	fmt.Println(students)

	//map的删除操作
	cities := map[string]string{"no1": "北京", "no2": "上海", "no3": "天津"}
	fmt.Println("cities = ", cities)
	delete(cities, "no1")
	fmt.Println("删除no1后,", cities)
	//当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println("删除no4后,", cities)

	//如果沃恩要删除map的所有key，没有一个专门的方法一次删除，可以遍历一下key，逐个删除；或者map = make(...)一个新的map，让原来的成为垃圾被GC掉

	//演示map的查找
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no1 key 值为%v\n", val)
	} else {
		fmt.Printf("没有no1 key\n")
	}

	//map的遍历只能使用for-range的结构来遍历
	for k, v := range cities {
		fmt.Printf("cities[%v] = %v\n", k, v)
	}

	//查看map的长度使用内置函数len(map)即可
	fmt.Println("students的map长度为:", len(students))
	for k1, v1 := range students {
		fmt.Println("k=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t %v = %v\n", k2, v2)
		}
	}

	//切片的数据类型如果是map，则我们称为 slice of map，map切片，这样使用则map个数就可以动态变化了
	//使用一个map来记录一个monster的信息name和age，也就是说一个monster对应一个map，并且妖怪的个数可以动态的增加=>map切片

	//1.声明一个map切片
	//var mapSlice []map[string]string
	//mapSlice = make([]map[string]string, 2) //准备放入两个妖怪[这里是先声明，再分配地址]
	var mapSlice []map[string]string = make([]map[string]string, 2) //准备放入两个妖怪[这里是声明+分配地址]
	//2.增减第一个妖怪的信息
	if mapSlice[0] == nil {
		mapSlice[0] = make(map[string]string, 2)
		mapSlice[0]["name"] = "玉兔精"
		mapSlice[0]["age"] = "200"
	}
	if mapSlice[1] == nil {
		mapSlice[1] = make(map[string]string, 2)
		mapSlice[1]["name"] = "牛魔王"
		mapSlice[1]["age"] = "400"
	}
	//这个写法越界了
	// if mapSlice[2] == nil {
	// 	mapSlice[2] = make(map[string]string, 2)
	// 	mapSlice[2]["name"] = "狐狸精"
	// 	mapSlice[2]["age"] = "300"
	// }

	//这里我们需要使用到切片的append函数，可以动态的增加妖怪
	//1.先定义一个monster信息
	newMap := map[string]string{"name": "狐狸精", "age": "200"}
	mapSlice = append(mapSlice, newMap)
	fmt.Println(mapSlice)

	//Go中灭有一个专门的方法针对map的key进行排序
	//Go中的map默认是无序的，注意也不是按照添加的顺序存放的，你每次遍历，得到的输出可能不一样
	//Go中map的排序，是先将key进行排序，然后根据key值遍历输出即可
	map1 := map[int]int{10: 20, 1: 10, 4: 15, 8: 90}
	fmt.Println("map1=", map1)

	//如果按照map的key的顺序进行排序输出
	//1.先将map的key放到切片中
	//2.对切片排序
	//3.遍历切片，然后按照key来输出map的值

	var keys []int
	for k := range map1 {
		keys = append(keys, k)
	}

	//排序
	sort.Ints(keys)

	fmt.Printf("keys=%v,len=%v,cap=%v\n", keys, len(keys), cap(keys))

	for _, k := range keys {
		fmt.Printf("map[%v]= %v\n", k, map1[k])
	}

	//map使用细节
	//1.map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map。
	modify(map1) //说明map是引用类型
	fmt.Println("map1 修改后的值 =", map1)
	//2.map的容量达到后，再想map增加元素，会自动扩容，并不会发生panic，也就是说map能动态的增长键值对。
	//3.map的value也经常使用struct类型，更适合管理复杂的数据（比前面value是一个map更好），比如value为Student结构体。

	//1)map的key为学生的学号，是唯一的；
	//2)map的value为结构体，包含学生的名字，年龄，地址

	studs := make(map[string]Student, 10)
	student1 := Student{"tom", 18, "北京"}
	student2 := Student{"mary", 28, "上海"}
	studs["no1"] = student1
	studs["no2"] = student2

	fmt.Println(students)

	//遍历各个学生信息
	for k, v := range studs {
		fmt.Printf("学生的编号是%v\n", k)
		fmt.Printf("学生的名字是%v\n", v.Name)
		fmt.Printf("学生的年龄是%v\n", v.Age)
		fmt.Printf("学生的地址是%v\n", v.Address)
		fmt.Println("==========================")
	}
}
