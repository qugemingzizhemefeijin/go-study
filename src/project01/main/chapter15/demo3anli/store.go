package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Monster 怪兽
type Monster struct {
	Name  string
	Age   int
	Skill string
}

//Store 序列化成JSON后保存到文件中
func (m *Monster) Store(filePath string) bool {
	content, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Store() Serial err =", err)
		return false
	}

	str := string(content)
	fmt.Println("monster serializable = ", str)

	//保存到文件中
	err = ioutil.WriteFile(filePath, content, 0666)
	if err != nil {
		fmt.Println("Store() Save err =", err)
		return false
	}
	return true
}

//ReStore 从文件中读取JSON并序列化成Monster结构体
func (m *Monster) ReStore(filePath string) bool {
	//首先判断读取文件正常
	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("File Not Exist Or Read Err =", err)
		return false
	}
	//如果没有问题，则读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read Error =", err)
		return false
	}

	str := string(content)
	fmt.Println("read from file content = ", str)

	//如果没有异常，则反序列化
	err = json.Unmarshal(content, m)
	if err != nil {
		fmt.Println("UnSerializable Error =", err)
		return false
	}

	return true
}
