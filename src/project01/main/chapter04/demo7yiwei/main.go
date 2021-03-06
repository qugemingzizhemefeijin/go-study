package main

import (
	"fmt"
)

//二进制转十进制规则：从最低位开始，将每个位上的数提取出来，乘以2的（位数-1）次方，然后求和。
//八进制转十进制规则，从最低位开始，将每个位上的数提取出来，乘以8的（位数-1）次方，然后求和。
//十六进制转十进制规则：上同

//十进制转二进制规则：将该数不断除以2，直到商为0为止，然后将每步得到的余数倒过来，就是对应的二进制。
//十进制转八进制规则：上同
//十进制转十六进制规则：上同

//二进制转八进制规则：将二进制数每三位一组（从低位开始组合），转成对应的八进制数即可。
//二进制转十六进制规则：将二进制数每四位一组（从低位开始组合），转成对应的十六进制数即可。

//八进制转二进制规则：将八进制数每一位转成对应的一个3位的二进制数即可。
//十六进制转二进制规则：将十六进制的每一位转成对应的一个4位的二进制数即可。
func main() {
	//右移运算符>>：低位溢出，符号位不变，并用符号位补溢出的高位
	//左移运算符<<：符号位不变，低位补0
	var a int = 1 >> 2 //0
	//1111 1111
	var b int = -1 >> 2 //
	var c int = 1 << 2  //4
	//1111 1111 -1补码
	//1111 1100	左移两位
	//1111 1011 算出反码
	//1000 0100 = 算出原码
	var d int = -1 << 2 //-4
	//a,b,c,d结果是多少
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("================")
	fmt.Println(2 & 3)  //2
	fmt.Println(2 | 3)  //3
	fmt.Println(13 & 7) //5
	fmt.Println(5 | 4)  //5
	//1000 0011	=> 1111 1100 => 1111 1101
	//							0000 0010
	//							1111 1111 补码
	//							1111 1110 反码
	//							1000 0001 原码 -1
	fmt.Println(-3 ^ 3)
	//fmt.Printf("%b", -1)
}
