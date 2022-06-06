package main

import "fmt"

func main() {
	//1. 假如还有97天放假，问：xx个星期零xx天
	var days int = 97
	fmt.Printf("%d天对应的是%d个星期零%d天\n", days, days/7, days%7)

	//定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为： 5/9*(华氏温度-100),请求出华氏温度对应的摄氏温度。
	var temperature float32 = 134.2
	//此处不能写为5/9，必须为5.0/9，否则5/9=0
	var normal float32 = 5.0 / 9 * (temperature - 100)
	fmt.Printf("华氏温度%f对应设置温度为：%f\n", temperature, normal)
}
