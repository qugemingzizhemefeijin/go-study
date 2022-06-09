package main

import (
	pkg "./pkga"
)

// 定义数值变量

// 如果出现 pkg_amd64.s:12: unexpected EOF，只需要在pkg_amd64.s最后一行换行即可
func main() {
	println(pkg.Id)
}
