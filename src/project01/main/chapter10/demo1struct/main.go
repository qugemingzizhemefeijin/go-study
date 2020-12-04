package main

import (
	"fmt"
)

//Cat 先定义一个cat结构体，将cat的各个字段信息，放到Cat结构体进行管理
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
	ptr   *int              //指针 默认都是nil
	slice []int             //切片 默认都是nil
	map1  map[string]string //map 默认都是nil
}

//Monster ss
type Monster struct {
	Name string
	Age  int
}

func main() {
	//张老太养了2只猫猫：一只名字叫小白，今年3岁，白色，还有一只叫小花，今年100岁，花色，
	//请编写一个程序，当用户输入小猫的名字时，就显示该猫的名字，年龄和颜色。如果用户输入的小猫名错误，则显示 张老太没有这只猫。

	//使用struct来完成案例

	//创建一个Cat的变量
	var cat1 Cat                      //var a int 非常像这种定义方式
	fmt.Printf("cat1的地址=%p\n", &cat1) //结构体是值类型
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃鱼"

	fmt.Println("cat1=", cat1)

	if cat1.ptr == nil {
		fmt.Println("ok1")
	}

	if cat1.slice == nil {
		fmt.Println("ok2")
	}

	if cat1.map1 == nil {
		fmt.Println("ok3")
	}

	//使用new内置函数，返回指针类型，默认new后指向的是零值
	cat1.ptr = new(int)
	*cat1.ptr = 999

	//使用slice，一定要先make再赋值等
	cat1.slice = make([]int, 10)
	cat1.slice[0] = 100

	//使用map，一定要先make map
	cat1.map1 = make(map[string]string)
	cat1.map1["key1"] = "abc"

	fmt.Println("cat1=", cat1)

	fmt.Println("猫猫的信息如下：")
	fmt.Println("name=", cat1.Name)
	fmt.Println("Age=", cat1.Age)
	fmt.Println("Color=", cat1.Color)
	fmt.Println("Hobby=", cat1.Hobby)
	fmt.Println("Ptr=", *cat1.ptr)

	//通过上面的案例和讲解我们可以看出：
	//1) 结构体是自定义的数据类型，代表一类事务
	//2) 结构体变量（实例）是具体的，实际的，代表一个具体变量

	//注意事项和细节说明
	//1) 在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值，规则同前面的一样，指针，slice，和map的零值都是nil，即还没有分配空间；
	//2) 不同结构体变量的字段是独立，互补影响，一个结构体变量字段的更改，不影响另外一个；(证明结构体是值拷贝)

	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1 //结构体是值类型，默认是值拷贝

	//这里两个输出是相同的
	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)
	fmt.Println("==============")
	//修改monster2的属性
	monster2.Name = "狐狸精"
	monster2.Age = 200
	//再次输出，monster1和monster2就不一样了，也就是证明两个指针指向的是不同的内存区域
	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)

	//创建结构体变量和访问结构体字段
	//1) 直接声明 var person Person
	//2) {}	var person Person = Person{Name:"123",...}	其中第一个Person是可以省略的
	var monster3 = Monster{Name: "山羊精", Age: 150}
	fmt.Println("monster3=", monster3)
	//3) & var person *Person = new(Person)
	var monster4 = new(Monster)
	(*monster4).Name = "琵琶精"
	//Go中实际上还可以这样子，直接用指针点出属性：原因是go的设计者，为了程序员使用方便，底层会对monster4.Age = 250 进行处理
	//会给 monster4 加上 取值运算(*monster4).Name = 250
	monster4.Age = 250
	fmt.Println("monster4=", *monster4)
	//4){} var person *Person = &Person{Name:"123",...}
	// var person = &Person{}
	var monster5 = &Monster{Name: "黑熊精", Age: 500}
	fmt.Println("monster5=", *monster5)

	var monster6 = &Monster{}
	monster6.Name = "白蛇精"
	monster6.Age = 800
	fmt.Println("monster6=", *monster6)
}
