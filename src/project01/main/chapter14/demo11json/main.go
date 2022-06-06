package main

import (
	"encoding/json"
	"fmt"
)

//Monster ... 注意此处使用了tag
type Monster struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Sal      float64
	Skill    string
}

func testStruct() {
	//演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-10",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
}

func testMap() {
	//定义一个map
	var a map[string]interface{}
	//使用map之前需要make
	a = make(map[string]interface{})

	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	//将a这个map进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	//输出序列化后的结果
	fmt.Printf("map序列化后=%v\n", string(data))
}

func testSlice() {
	//定义一个slice
	var slice []map[string]interface{}

	m1 := make(map[string]interface{})
	m1["name"] = "铁扇公主"
	m1["age"] = 20
	m1["address"] = [2]string{"风云洞", "牛魔洞"}

	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "白骨精"
	m2["age"] = 100
	m2["address"] = "白骨洞"

	slice = append(slice, m2)

	//将a这个map进行序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	//输出序列化后的结果
	fmt.Printf("切片序列化后=%v\n", string(data))
}

//对基本类型数据序列化，对基本数据类型进行序列化意义不大，也无法解析成JSON
func testFloat64() {
	num1 := 2345.67

	//将a这个map进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	//输出序列化后的结果
	fmt.Printf("对Float序列化后=%v\n", string(data))
}

//结构、切片和map的序列化案例
func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
