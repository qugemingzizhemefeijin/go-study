package main

import (
	"log"
	"math/rand"
)

var endpoints = []string{
	"100.69.62.1:3232",
	"100.69.62.32:3232",
	"100.69.62.42:3232",
	"100.69.62.81:3232",
	"100.69.62.11:3232",
	"100.69.62.113:3232",
	"100.69.62.101:3232",
}


// 修正洗牌算法
// 从数学上得到过证明的还是经典的 fisher-yates 算法，主要思路为每次随机挑选一个值，放在数组末尾。
// 然后在 n-1 个元素的数组中再随机挑选一个值，放在数组末尾，以此类推。
//func shuffle(indexes []int) {
//	for i:=len(indexes); i>0; i-- {
//		lastIdx := i - 1
//		idx := rand.Intn(i)
//		indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
//	}
//}

// 在 Go 的标准库中已经为我们内置了该算法:
func shuffle(n int) []int {
	b := rand.Perm(n)
	return b
}

func request(params map[string]interface{}) error {
	var err error

	indexes := shuffle(len(endpoints))
	maxRetryTimes := 3

	idx := 0
	for i := 0; i < maxRetryTimes; i++ {
		err = apiRequest(params, indexes[idx])
		if err == nil {
			break
		}
		idx++
	}

	if err != nil {
		// logging
		return err
	}
	return nil
}

func apiRequest(params map[string]interface{}, idx int) error {
	log.Println("idx = ", idx, ", point = ", endpoints[idx])
	return nil
}

// 段简短的程序里有两个隐藏的隐患:
// 1.没有随机种子。在没有随机种子的情况下，rand.Intn() 返回的伪随机数序列是固定的。
// 2.洗牌不均匀，会导致整个数组第一个节点有大概率被选中，并且多个节点的负载分布不均衡。
func main() {
	var params map[string]interface{}

	for i := 0; i < 3; i++ {
		request(params)
	}
}
