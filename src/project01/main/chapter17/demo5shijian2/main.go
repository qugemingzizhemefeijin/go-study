package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//Monster 这里申明了一个Monster结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

//Print ...
func (m Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(m)
	fmt.Println("---end---")
}

//TestStruct ...
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	fmt.Printf("type = %v, kind=%v\n", typ, kd)
	if kd != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	fmt.Printf("elem kind = %v\n", val.Elem().Kind())

	fmt.Println("========================")
	num := val.Elem().NumField()
	//Value Field
	val.Elem().Field(0).SetString("白象精")
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v %v = %v \n", i, val.Elem().Field(i).Kind(), typ.Elem().Field(i).Name, val.Elem().Field(i))
	}

	fmt.Printf("struct has %d fields\n", num)
	fmt.Println("========================")

	//Type StructField
	tag := typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//调用Print()方法
	val.Elem().Method(0).Call(nil)
}

//使用反射的方式来获取结构体的tag标签，遍历字段的值，修改字段值，调用结构体方法等，要求通过传递地址的方式完成
func main() {
	m := Monster{
		Name:  "黄鼠狼精",
		Age:   500,
		Score: 98.5,
		Sex:   "女",
	}
	//先说明一下，Marshal就是通过反射获取到struct的tag值的
	result, _ := json.Marshal(m)
	fmt.Println("json result:", string(result))

	TestStruct(&m)
	fmt.Println("exit success")
}
