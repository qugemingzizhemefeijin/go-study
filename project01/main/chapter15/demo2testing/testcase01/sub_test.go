package test

import (
	"testing" //引入go的testing测试包
)

func TestGetSub(t *testing.T) {
	//调用
	res := getSub(10, 3)
	if res != 7 {
		t.Fatalf("getSub(10, 3)执行错误，期望值=%v 实际值=%v\n", 7, res)
	}

	//如果正确，输出日志
	t.Logf("getSub(10, 3) 执行正确...")
}
