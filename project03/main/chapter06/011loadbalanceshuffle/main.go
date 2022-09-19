package main

import (
	"log"
	"math/rand"
)

// 考虑到我们需要随机选取每次发送请求的节点， 同时在遇到下游返回错误时换其它节点重试。 所以我们设计⼀个⼤⼩和节点数组⼤⼩⼀致的索引数组，
// 每次来新的请求， 我们对索引数组做洗牌， 然后取第⼀个元素作为选中的服务节点， 如果请求失败， 那么选择下⼀个节点重试， 以此类推。

var endpoints = []string{
	"100.69.62.1:3232",
	"100.69.62.32:3232",
	"100.69.62.42:3232",
	"100.69.62.81:3232",
	"100.69.62.11:3232",
	"100.69.62.113:3232",
	"100.69.62.101:3232",
}

func shuffle(slice []int) {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
}

func request(params map[string]interface{}) error {
	var indexes = []int{0, 1, 2, 3, 4, 5, 6}
	var err error

	shuffle(indexes)
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
