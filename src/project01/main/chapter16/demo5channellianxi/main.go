package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//按理：
//1. 创建一个Person结构体 [Name, Age, Address]
//2. 使用rand方法配合随机创建10个Person实例，并放入到channel中
//3. 遍历channel，将各个Person实例的信息显示在终端

//Person ...
type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	personChan := make(chan Person, 10)
	//这里放入10个Person到管道中
	for i := 0; i < 10; i++ {
		//设置随机种子(这里也必须暂时一下，不然太快了，随机种子都是一样的)
		time.Sleep(time.Nanosecond * 10)
		rand.Seed(time.Now().UnixNano())
		//随机生成一些信息
		personChan <- Person{
			Name:    "橙子-" + strconv.Itoa(rand.Intn(1000)+100),
			Age:     rand.Intn(100) + 1,
			Address: "北京-" + strconv.Itoa(rand.Intn(10000)+1000),
		}
	}
	//time.Sleep(time.Second * 10)
	fmt.Printf("personChan len=%d, cap=%d \n", len(personChan), cap(personChan))
	//这里循环去除Person
	len := len(personChan)
	for i := 0; i < len; i++ {
		fmt.Printf("Person[%d] = %v \n", i+1, <-personChan)
	}
}
