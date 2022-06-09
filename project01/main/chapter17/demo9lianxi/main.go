package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//Cal ...
type Cal struct {
	Num1, Num2 int
}

//GetSub ...
func (c Cal) GetSub(name string) string {
	res := c.Num1 - c.Num2
	return name + " 完成了减法运行，" + strconv.Itoa(c.Num1) + " - " + strconv.Itoa(c.Num2) + " = " + strconv.Itoa(res)
}

//练习：
//1. 编写一个Cal结构体，有两个字段 Num1和Num2
//2. 方法GetSub(name string)
//3. 使用反射遍历Cal结构体所有的字段信息
//4. 使用反射机制完成对GetSub 的调用，输出形式为："tom 完成了减法运行，8-3=5"
func main() {
	c := &Cal{
		Num1: 8,
		Num2: 5,
	}

	rType := reflect.TypeOf(c)   //这里得到的是*Cal
	rValue := reflect.ValueOf(c) //同理
	fmt.Printf("rType=%v, rValue=%v\n", rType.Kind(), rValue.Kind())

	//此处需要拿到真正的结构信息
	rType = rType.Elem()
	rValue = rValue.Elem()
	fmt.Printf("rType=%v, rValue=%v\n", rType.Kind(), rValue.Kind())

	//获取所有的字段数量以及遍历获取其信息和值
	fieldNum := rValue.NumField()
	fmt.Printf("field num %d \n", fieldNum)

	for i := 0; i < fieldNum; i++ {
		fmt.Printf("%d %v %v = %v\n", i, rValue.Field(i).Kind(), rType.Field(i).Name, rValue.Field(i).Interface())
	}

	//再通过反射调用GetSub方法
	numOfMethod := rValue.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//获取方法调用的参数数量
	method := rType.Method(0)
	numOfParam := rValue.Method(0).Type().NumIn()
	fmt.Printf("struct method %v , param count = %d\n", method.Name, numOfParam)

	//初始化调用参数
	args := []reflect.Value{reflect.ValueOf("jerry")}

	res := rValue.Method(0).Call(args)

	fmt.Println(res[0])
	fmt.Println("exit success")
}
