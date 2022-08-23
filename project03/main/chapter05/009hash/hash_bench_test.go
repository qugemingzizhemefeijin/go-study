package main

import "testing"

// go test -bench=.
// go test -bench=.  -cpu=1,2,4
// go test -bench=Fib320 -count=10

// 基准测试
// https://blog.csdn.net/qq_28119741/article/details/117935237

func BenchmarkMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5Hash()
	}
}

func BenchmarkSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha1Hash()
	}
}

func BenchmarkMurmurHash32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur32()
	}
}

func BenchmarkMurmurHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur64()
	}
}
