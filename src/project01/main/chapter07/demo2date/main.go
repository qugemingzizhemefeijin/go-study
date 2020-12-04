package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	//时间和日期相关的函数
	//1.获取当前时间
	now := time.Now()
	fmt.Printf("now=%v, now type = %T\n", now, now)
	//2.如何获取其他的时间信息，获取年月日，时分秒
	fmt.Printf("当前日期：%d-%d-%d %d:%d:%d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	//3. 如何格式化日期
	// 一:
	dateStr := fmt.Sprintf("%d-%d-%d %d:%d:%d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("dateStr = ", dateStr)
	// 二： 2006/01/02 15:04:05 这个字符串的各个数字是固定的，必须要这么写
	fmt.Printf("日期时间：%v \n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("日期：%v \n", now.Format("2006-01-02"))
	fmt.Printf("时间：%v \n", now.Format("15:04:05"))

	//4.时间常量
	//Nanosecond 纳秒 Microsecond 微秒 Millisecond 毫秒 Second 秒 Minute 分 Hour 时
	//fmt.Println("Second=", time.Duration(10)*time.Second)
	// fmt.Println("Second=", int64(time.Second/time.Millisecond))
	// fmt.Println("start...")
	// time.Sleep(1 * time.Second) //暂停1秒
	// fmt.Println("ending...")

	//需求：每隔0.1秒打印一个数字，打印到100时就退出
	// i := 0
	// for {
	// 	i++
	// 	fmt.Println(i)
	// 	//休眠
	// 	time.Sleep(time.Millisecond * 100)
	// 	if i == 100 {
	// 		break
	// 	}
	// }

	//5. 获取当前unix时间戳和unixnano时间戳（作用是可以获取随机数字）
	fmt.Println("unix time = ", now.Unix())
	fmt.Println("unix nano time = ", now.UnixNano())

	//统计函数的执行时间
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("test03 总计执行的秒数：%d\t", end-start)
}
