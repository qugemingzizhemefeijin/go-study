package main

import (
	"fmt"
	"github.com/murmur3"
)

// 对于哈希算法来说，除了性能方面的问题，还要考虑哈希后的值是否分布均匀。如果哈希后的值分布不均匀，那也自然就起不到均匀灰度的效果了。
// 以 murmur3 为例，我们先以 15810000000 开头，造一千万个和手机号类似的数字，然后将计算后的哈希值分十个桶，并观察计数是否均匀。

func murmur64(p string) uint64 {
	return murmur3.Sum64([]byte(p))
}

var bucketSize = 10

func main() {
	var bucketMap = map[uint64]int{}
	for i:= 15000000000; i< 15000000000 + 10000000; i++ {
		hashInt := murmur64(fmt.Sprint(i)) % uint64(bucketSize)
		bucketMap[hashInt]++
	}
	fmt.Println(bucketMap)
}
