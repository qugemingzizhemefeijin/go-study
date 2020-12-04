package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

//1.开一个协程writeDataToFile，随机生成1000个数据，存放到文件中
//2.档writeDataToFile完成写1000个数据到文件后，让sort协程从文件中读取1000个数据，并完成排序，重新写入到另外一个文件
//3.考察点：协程和管道+文件的综合使用
//4.功能扩展：开10个协程writeDataToFile，每个协程随机生成1000个数据，存放到10个文件中
//5.档10个文件都生成了，让10个sort协程从10个文件中读取1000个数据，并完成排序，重新写入到10个结果文件

func writeDataToFile(i int, wait *sync.WaitGroup) {
	filePath := "E:/" + fmt.Sprintf("%d", i) + ".txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file %v error=%v \n", filePath, err)
		return
	}
	//关闭文件
	defer file.Close()
	//打开文件缓冲
	writer := bufio.NewWriter(file)
	//循环写入
	for i := 0; i < 1000; i++ {
		//设置种子
		rand.Seed(time.Now().UnixNano())
		//随机生成0-10000之内的数字
		writer.WriteString(fmt.Sprintf("%d", rand.Intn(10000)) + "\n")
		time.Sleep(time.Nanosecond)
	}

	writer.Flush()

	wait.Done()
}

type intSlice []int

func (p intSlice) Len() int {
	return len(p)
}

func (p intSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p intSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortFromFile(i int, wait *sync.WaitGroup) {
	filePath := "E:/" + fmt.Sprintf("%d", i) + ".txt"
	//这里就可以读取数据到int切片中了
	var data intSlice

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("open file %v error=%v \n", filePath, err)
		return
	}
	//关闭文件
	defer file.Close()
	//打开读取文件缓冲
	reader := bufio.NewReader(file)
	for {
		str, _, err := reader.ReadLine()
		if err != nil || len(str) == 0 {
			break
		}
		a, _ := strconv.ParseInt(string(str), 10, 64)
		data = append(data, int(a))
	}
	//读取完毕则排序
	sort.Sort(data)

	//排序完毕后，重新写入另一个文件
	idx := strings.LastIndex(filePath, ".")
	//新文件路径
	newFilePath := filePath[:idx] + "_new." + filePath[idx+1:]

	writeDataToNewFile(newFilePath, data)

	wait.Done()
}

//写入到新文件中
func writeDataToNewFile(filePath string, data intSlice) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open new file %v error=%v \n", filePath, err)
		return
	}

	//关闭文件
	defer file.Close()
	//打开文件缓冲
	writer := bufio.NewWriter(file)

	for i := 0; i < len(data); i++ {
		writer.WriteString(fmt.Sprintf("%d", data[i]) + "\n")
	}
	//刷新缓冲区
	writer.Flush()
}

func main() {
	//生成一个信号锁
	var wait sync.WaitGroup

	wait.Add(10)
	//开启协程写入文件
	for i := 0; i < 10; i++ {
		go writeDataToFile(i, &wait)
	}
	//主线程等待通知完成
	wait.Wait()

	wait.Add(10)
	//开启10个协程进行排序
	for i := 0; i < 10; i++ {
		go sortFromFile(i, &wait)
	}
	//主线程等待通知完成
	wait.Wait()

	fmt.Println("exit success")
}
