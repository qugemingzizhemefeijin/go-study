package main

import (
	"encoding/gob"
	"os"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

// 在 go 语言中，这个过程使用 encoding/gob 完成，gob 对 go 就像 Serialization 对 java、pickle 对 python，
// 这些序列化方案都是语言内部的，需要跨语言的描述就需要 xml、json、protocol buffers 序列化方案了。
func main() {
	p1 := Person{"zhangsan", 20}
	enc1 := gob.NewEncoder(os.Stdout)
	_ = enc1.Encode(p1)
	_, _ = os.Stdout.WriteString("\n")

	// 这里当然也可以保存到文件中
	file1, _ := os.OpenFile("E:/person.gob", os.O_CREATE | os.O_WRONLY, 0644)
	enc2 := gob.NewEncoder(file1)
	_ = enc2.Encode(p1)
	_ = file1.Close()

	// 再从文件中读取数据
	var p2 Person
	file2, _ := os.Open("E:/person.gob")
	defer file2.Close()
	dec := gob.NewDecoder(file2)
	_ = dec.Decode(&p2)

	_, _ = os.Stdout.WriteString("Name:" + p2.Name + ",Age:" + strconv.Itoa(p2.Age) + "\n")
}
