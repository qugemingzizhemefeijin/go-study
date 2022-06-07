package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func MyRecover() interface{} {
	log.Println("trace...")
	return recover()
}

// 当希望将捕获到的异常转为错误时， 如果希望忠实返回原始的信息， 需要针对不同的类型分别处理：
func foo() (err error) {
	defer func() {
		if r:= recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("Unknown panic: %v", r)
			}
		}
	}()

	panic("TODO")
}

// 必须要和有异常的栈帧只隔⼀个栈帧， recover 函数才能正常捕获异常。 换⾔之， recover 函数捕获的是祖⽗⼀级调⽤函数栈帧的异常（刚好可以跨越⼀层 defer 函数）
func main() {
	// ⽐如， 有时候我们可能希望包装⾃⼰的 MyRecover 函数，在内部增加必要的⽇志信息然后再调⽤ recover，这是错误的做法
	defer func() {
		// 无法捕获异常
		if r:= MyRecover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic(1)

	// 同样， 如果是在嵌套的 defer 函数中调⽤ recover 也将导致⽆法捕获异常：

	// 2层嵌套的 defer 函数中直接调⽤ recover 和1层 defer 函数中调⽤包装的 MyRecover 函数⼀样， 都是经过了2个函数帧才到达真正的 recover 函数，
	// 这个时候Goroutine的对应上⼀级栈帧中已经没有异常信息
	defer func() {
		defer func() {
			// 无法捕获异常
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
	}()

	panic(1)

	// 如果我们直接在 defer 语句中调⽤ MyRecover 函数⼜可以正常⼯作了：
	defer MyRecover()
	panic(1)

	// ⽆法捕获异常
	defer recover()
	panic(1)

	// 为了避免 recover 调⽤者不能识别捕获到的异常, 应该避免⽤ nil 为参数抛出异常
	defer func() {
		if r:= recover(); r != nil {

		}
		// 虽然总是返回nil, 但是可以恢复异常状态
	}()

	// 警告: ⽤`nil`为参数抛出异常
	panic(nil)
}
