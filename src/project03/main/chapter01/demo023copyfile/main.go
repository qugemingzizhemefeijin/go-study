package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	i, _ := CopyFile("E:/Test1.class", "E:/Test.class")
	fmt.Println(i)
}
