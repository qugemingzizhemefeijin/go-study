package main

import (
	"bytes"
	"fmt"
	"image/color"
	"io"
	"os"
)

// 方法是由函数演变过来的，只是将函数的第一个对象参数移动到了函数名前面而已。我们也可以直接当做普通的函数来使用
// var CloseFile = (*File).Close
// var ReadFile = (*File).Read
// f. _ := OpenFile("foo.dat")
// ReadFile(f, 0, data)
// CloseFile(f)

// 继承

// Point desc
type Point struct{ X, Y float64 }

// ColoredPoint desc
type ColoredPoint struct {
	Point
	Color color.RGBA
}

// UpperWriter desc
type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	fmt.Println("===============")
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
}
