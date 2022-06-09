package main

// int sum(int a, int b) { return a+b; }
import "C"

// 此命令可以查看CGO生成的中间文件 go tool cgo main.go
// 查看文件信息 ls _obj/ |awk '{print $NF}'
/*
_cgo_.o
_cgo_export.c
_cgo_export.h
_cgo_flags
_cgo_gotypes.go
_cgo_main.c
main.cgo1.go
main.cgo2.c

其中 _cgo_.o 、 _cgo_flags 和 _cgo_main.c ⽂件和我们的代码没有直接的逻辑关联
 */
func main() {
	println(C.sum(1, 1))
}
