package utils

import (
	"fmt"
)

//Calc 将计算的功能，放到一个函数中，在需要的地方调用即可
func Calc(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误...")
	}
	return res
}
