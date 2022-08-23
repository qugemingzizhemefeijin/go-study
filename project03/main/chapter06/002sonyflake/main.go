package main

import (
	"fmt"
	"github.com/sonyflake"
	"os"
	"time"
)

// sonyflake 是 Sony 公司的一个开源项目，基本思路和 snowflake 差不多，不过位分配上稍有不同
// 1B Unused | 39B Timestamp | 8B Sequence ID | 16B Machine ID |
// 这里的时间只用了 39 个 bit，但时间的单位变成了 10ms，所以理论上比 41 位表示的时间还要久 (174 年)。
// Sequence ID 和之前的定义一致，Machine ID 其实就是节点 id。sonyflake 与众不同的地方在于其在启动阶段的配置参数。
/*
func NewSonyflake(st Settings) *Sonyflake

type Settings struct {
	StartTime      time.Time
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

StartTime 选项和我们之前的 Epoch 差不多，如果不设置的话，默认是从 2014-09-01 00:00:00 +0000 UTC 开始。

MachineID 可以由用户自定义的函数，如果用户不定义的话，会默认将本机 IP 的低 16 位作为 machine id。

CheckMachineID 是由用户提供的检查 MachineID 是否冲突的函数。这里的设计还是比较巧妙的，
如果有另外的中心化存储并支持检查重复的存储，那我们就可以按照自己的想法随意定制这个检查 MachineID 是否冲突的逻辑。
如果公司有现成的 Redis 集群，那么我们可以很轻松地用 Redis 的集合类型来检查冲突。
 */

func getMachineID() (uint16, error) {
	var machineID uint16
	var err error
	machineID = readMachineIDFromLocalFile()
	if machineID == 0 {
		machineID, err = generateMachineID()
		if err != nil {
			return 0, err
		}
	}

	return machineID, nil
}

func saddMachineIDToRedisSet() (int, error) {
	return 0, nil
}

func generateMachineID() (uint16, error) {
	return 1, nil
}

func readMachineIDFromLocalFile() uint16 {
	return 0
}

func saveMachineIDToLocalFile(machineID uint16) error {
	fmt.Println("保存machineID", machineID)
	return nil
}

func checkMachineID(machineID uint16) bool {
	saddResult, err := saddMachineIDToRedisSet()
	if err != nil || saddResult == 0 {
		return true
	}

	err = saveMachineIDToLocalFile(machineID)
	if err != nil {
		return true
	}

	return false
}

func main() {
	t, _ := time.Parse("2006-01-02", "2022-01-01")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
		CheckMachineID: checkMachineID,
	}

	sf := sonyflake.NewSonyflake(settings)
	if sf == nil {
		panic("sonyflake not created")
	}
	id, err := sf.NextID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("ID: ", id)
}