package main

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"
)


// 获取goid偏移值，其实是自己维护的值
var offsetDictMap = map[string]int64{
	"go1.14": 152,
	"go1.10": 152,
	"go1.9":  152,
	"go1.8":  192,
}

var g_goid_offset = func() int64 {
	goversion := runtime.Version()
	fmt.Println(goversion)
	for key, off := range offsetDictMap {
		if goversion == key || strings.HasPrefix(goversion, key) {
			return off
		}
	}
	panic("unsupported go version:"+goversion)
}()

func GetGroutineId() int64 {
	g := getg()
	p := (*int64)(unsafe.Pointer(uintptr(g) + g_goid_offset))
	return *p
}

func main() {
	// 没搞通，o(╥﹏╥)o
	fmt.Println(GetGroutineId())
}
