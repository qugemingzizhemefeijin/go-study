package main

import (
	"fmt"
)

func main() {
	//使用switch把小写类型的char型转为大写（键盘输入）。只转换a,b,c,d,e，其它的输出'other'
	// var c byte
	// fmt.Println("请输入字符:")
	// fmt.Scanf("%c", &c)
	// switch c {
	// // case 'a':
	// // 	fmt.Println("A")
	// // case 'b':
	// // 	fmt.Println("B")
	// // case 'c':
	// // 	fmt.Println("C")
	// // case 'd':
	// // 	fmt.Println("D")
	// // case 'e':
	// // 	fmt.Println("E")
	// case 'a', 'b', 'c', 'd', 'e':
	// 	fmt.Printf("%c", c-32)
	// default:
	// 	fmt.Println("other")
	// }

	//对学生成绩大于60分的，输出“合格”。低于60分的，输出“不合格”。（注：输入的成绩不能大于100）
	// var score byte
	// fmt.Println("请输入成绩")
	// fmt.Scanln(&score)
	// switch {
	// case score >= 60 && score <= 100:
	// 	fmt.Println("合格")
	// case score <= 60:
	// 	fmt.Println("不合格")
	// default:
	// 	fmt.Println("错误")
	// }

	//根据用户指定月份，打印该月份所属的季节。3,4,5春季，6,7,8夏季，9,10,11秋季,12,1,2冬季
	// var month byte
	// fmt.Println("请输入月份")
	// fmt.Scanln(&month)
	// switch month {
	// case 3, 4, 5:
	// 	fmt.Println("春季")
	// case 6, 7, 8:
	// 	fmt.Println("夏季")
	// case 9, 10, 11:
	// 	fmt.Println("秋季")
	// case 12, 1, 2:
	// 	fmt.Println("冬季")
	// default:
	// 	fmt.Println("输入错误")
	// }

	//根据用户输入显示对应的星期时间(string)，如果“星期一”，显示“干煸豆角”，如果“星期二”显示“醋溜土豆”
	//星期三“红烧狮子头”，星期四“油炸花生米”，星期五“蒜蓉扇贝”，星球六“东北乱炖”，星期日“大盘鸡”
	var week string
	fmt.Println("请输入星期日期:")
	_, _ = fmt.Scanln(&week)
	switch week {
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("醋溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("油炸花生米")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期日" , "星期天":
		fmt.Println("大盘鸡")
	default:
		fmt.Println("输入错误")
	}
}
