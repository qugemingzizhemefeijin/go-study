package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

type StudentSlice []Student

func (s StudentSlice) Len() int {
	return len(s)
}

func (s StudentSlice) Less(i, j int) bool {
	return s[i].Score > s[j].Score //从大到小排序
}

func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var stus = StudentSlice{
		Student{Name: "小明", Age: 20, Score: 78.8},
		Student{Name: "小红", Age: 22, Score: 98.8},
		Student{Name: "小花", Age: 28, Score: 18.8},
		Student{Name: "小刚", Age: 21, Score: 28.8},
		Student{Name: "二蛋", Age: 18, Score: 68.8},
	}

	sort.Sort(stus)

	for i := 0; i < len(stus); i++ {
		fmt.Printf("Name=%v,Age=%v,Score=%v\n", stus[i].Name, stus[i].Age, stus[i].Score)
	}
}
