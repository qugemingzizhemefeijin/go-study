package main

import (
	"fmt"
	"github.com/snowflake"
	"os"
)

// https://github.com/bwmarrin/snowflake 是一个相当轻量化的 snowflake 的 Go 实现。
// 1B Unused | 41B Timestamp | 10B NodeID | 12B Sequence ID |
// 这个库也给我们留好了定制的后路，其中预留了一些可定制字段：
// Epoch int64 = 1288834974657   起始时间
// Node Bits uint8 = 10           机器编号的位长 （Node + Step 最多不能超过22位）
// Step Bits uint8 = 12           自增序列的位长
func main() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println(
			"Id: ", id,
			"Node: ", id.Node(),
			"Step: ", id.Step(),
			"Time: ", id.Time(),
		)
	}
}