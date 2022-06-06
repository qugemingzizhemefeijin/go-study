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

func unMarshalStruct() {
	str := "{\"name\":\"牛魔王\",\"age\":500,\"birthday\":\"2011-11-10\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	//定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster) //注意：此处monster是要传递一个指针
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
		return
	}

	fmt.Printf("反序列化后结构：%v\n", monster)
}

func unMarshalMap() {
	str := "{\"address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}"

	//定义map
	var m map[string]interface{}
	//反序列化不需要make，因为make操作被封装到了Unmarshal函数中了
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
		return
	}

	fmt.Printf("反序列化后Map：%v\n", m)
}

func unMarshalSlice() {
	str := "[{\"address\":[\"风云洞\",\"牛魔洞\"],\"age\":20,\"name\":\"铁扇公主\"},{\"address\":\"白骨洞\",\"age\":100,\"name\":\"白骨精\"}]"

	//定义slice
	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
		return
	}

	fmt.Printf("反序列化后Slice：%v\n", slice)
}

func main() {
	//将json字符串反序列化成结构体、map和切片
	unMarshalStruct()
	unMarshalMap()
	unMarshalSlice()
}
