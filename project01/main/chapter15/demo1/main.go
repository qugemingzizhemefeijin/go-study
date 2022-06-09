package main

import (
	"fmt"
)

//一个被测试的函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

//1.不方便，我们需要在main函数中调用，这样就需要去修改main函数，如果现在项目正在运行，就可能去停止项目
//2.不利于管理，因为当我们测试多个函数或者多个模块时，都需要写在main函数中，不利于我们管理和清晰我们思路
//3.引出单元测试 testing 测试框架 可以很好的解决上面的问题
func main() {
	//传统的测试方法，就是在main函数中使用看看结果是否正确
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper错误 返回值=%v 期望值=%v", res, 55)
	} else {
		fmt.Printf("addUpper正确 返回值=%v 期望值=%v", res, 55)
	}
}
