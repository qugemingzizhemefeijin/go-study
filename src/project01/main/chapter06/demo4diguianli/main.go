package main

import (
	"fmt"
)

//1.斐波那契数列，给出N，求出第N个斐波那契数列是多少？
//1,1,2,3,5,8,13....
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fbn(n-1) + fbn(n-2)
}

//2.已知f(1)=3;f(n) = 2*f(n-1)+1;请使用递归的思想，求出f(n)的值
func f(n int) int {
	if n == 1 {
		return 3
	}
	return 2*f(n-1) + 1
}

//3.有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个！以后每天猴子吃其中的一半，然后再多吃一个。
//  当到第十天时，想再吃时（还没吃），发现只有1个桃子了。问：最初共多少个桃子？
func eatPeach(day int) int {
	if day > 10 || day < 1 {
		fmt.Println("输入的天数不对")
		return 0
	}
	if day == 10 {
		return 1
	}
	return (eatPeach(day+1) + 1) * 2
}

func main() {
	fmt.Println("res=", fbn(4))

	fmt.Println("f(1)=", f(1))
	fmt.Println("f(5)=", f(5))

	fmt.Printf("第一天的桃子数%d\n", eatPeach(1))
}
