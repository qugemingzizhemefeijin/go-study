package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string = "Hello北"
	fmt.Println(s)

	//1.统计字符串的长度，按字节len(str)
	fmt.Println("1. str len = ", len(s))

	//2.字符串遍历，同时处理有中文的问题r:=[]rune(str)
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c ", r[i])
	}
	fmt.Println()

	//3.字符串转整数,n,err := strconv.Atoi("12")
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Printf("3. atoi 转换错误 \n")
	} else {
		fmt.Printf("3. atoi = %v, Type %T\n", n, n)
	}

	//4.整数转字符串,str = strconv.Itoa(12345)
	s = strconv.Itoa(n)
	fmt.Printf("4. itoa = %v, Type %T\n", s, s)

	//5.字符串转[]byte: var bytes = []byte("hello go")
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		fmt.Printf("%c ", b[i])
	}
	fmt.Println()

	//6.[]byte转字符串: str = string([]byte{97,98,99})
	s = string([]byte{97, 98, 99})
	fmt.Printf("6. []byte to string, %v \n", s)

	//7.10进制转2,8,16进制：str = strconv.FormatInt(123,2)// 2->8,16
	s = strconv.FormatInt(10, 2)
	fmt.Println("7. 十进制转二进制 ", s)

	//8.查找字串是否在指定的字符串中：strings.Contains("seafood", "foo")//true
	s = "seafood"
	if strings.Contains(s, "foo") {
		fmt.Println("8. 字符串包含foo")
	} else {
		fmt.Println("8. 字符串不包含foo")
	}

	//9.统计一个字符串有几个指定的字串：strings.Count("ceheese", "e")
	s = "ceheese"
	fmt.Printf("9. 字符串 %v 包含 %d 个 e\n", s, strings.Count(s, "e"))

	//10.不区分大小写的字符串比较:fmt.Println(strings.EqualFold("abc", "ABC"))//true
	s = "abc"
	fmt.Printf("10. 字符串 不区分大小写比较 %t\n", strings.EqualFold(s, "Abc"))

	//11.返回字串在字符串第一次出现的Index值，如果没有返回-1: strings.Index("INT_abc", "abc") //4
	s = "INT_abc"
	s2 := "abc"
	idx := strings.Index(s, s2)
	if idx == -1 {
		fmt.Printf("11. 字符串 %v 不包含 %v \n", s, s2)
	} else {
		fmt.Printf("11. 字符串 %v 包含 %v 的索引在 %d\n", s, s2, idx)
	}

	//12.返回字串在字符串最后一次出现的Index值，如没有返回-1：strings.LastIndex("go golang", "go")
	idx = strings.LastIndex(s, s2)
	if idx == -1 {
		fmt.Printf("12. 字符串 %v 没有 %v \n", s, s2)
	} else {
		fmt.Printf("12. 字符串 %v 最后一次出现 %v 的索引在 %d\n", s, s2, idx)
	}

	//13.将指定的字串替换成 另外一个字串：strings.Replace("go go hello", "go", "go语言", n) n可以指定你希望替换几个，如果n=-1表示全部替换
	s = strings.Replace("go go hello", "go", "go语言", -1)
	fmt.Printf("13. 替换 %v \n", s)

	//14.按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello world,ok", ",")
	sarry := strings.Split("go go hello", " ")
	for i := 0; i < len(sarry); i++ {
		fmt.Println(sarry[i])
	}

	//15.将字符串的字母进行大小写的转换：strings.ToLower("Go")
	fmt.Println("15. =====>", strings.ToLower("Go"), "===", strings.ToUpper("Go"))

	//16.将字符串左右两边的空格去掉：strings.TrimSpace(" tn a lone gopher ntrn ")
	fmt.Printf("16. 去除左右空格 ===>%v\n", strings.TrimSpace(" tn a lone gopher ntrn "))

	//17.将字符串左右两边指定的字符去掉：strings.Trim("! Hello !", " !") //["hello"] 将左右两边!和" "去掉
	fmt.Printf("17. 去除左右指定字符 ===>%v\n", strings.Trim("! Hello !", " !"))

	//18.将字符串左边指定的字符去掉：strings.TrimLeft("! hello !", " !") //将左边!和" "去掉
	fmt.Printf("18. 去除字符串左边指定的字符 ===>%v\n", strings.TrimLeft("! hello !", " !"))

	//19.将字符串右边指定的字符去掉：strings.TrimRight("! hello !", " !") //将右边!和" "去掉
	fmt.Printf("19. 去除字符串右边指定的字符 ===>%v\n", strings.TrimRight("! Hello !", " !"))

	//20.判断字符串是否以指定的字符串开头：strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("20. 判断字符串是否以指定的字符串开头 ===> %v \n", strings.HasPrefix("https://www.baidu.com", "https"))

	//21.判断字符串是否以指定的字符串结尾：string.HasSuffix("NLT_abc.jpg", "abc")
	fmt.Printf("21. 判断字符串是否以指定的字符串结尾 ===> %v \n", strings.HasSuffix("E:/1.jpg", ".jpg"))
}
