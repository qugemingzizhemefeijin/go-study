package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//AddMonster 录入三个妖怪
func AddMonster(conn redis.Conn) {
	var name string
	var age int
	var skill string

	for i := 1; i <= 3; i++ {
		fmt.Println("请输入妖怪", i, "名称：")
		fmt.Scanln(&name)
		fmt.Println("请输入妖怪", i, "年龄：")
		fmt.Scanln(&age)
		fmt.Println("请输入妖怪", i, "技能：")
		fmt.Scanln(&skill)

		//添加到redis中
		_, err := conn.Do("HMSet", "monster"+fmt.Sprintf("%d", i), "name", name, "age", age, "skill", skill)
		if err != nil {
			fmt.Printf("add monster %d err = %v \n", i, err)
		}
	}
}

//TraversalMonster 遍历所有的妖怪
func TraversalMonster(conn redis.Conn) {
	//先取出所有的monster开头的keys
	keys, err := redis.Strings(conn.Do("Keys", "monster*"))
	if err != nil {
		fmt.Println("Keys err", err)
		return
	}
	//如果成功，则循环keys
	for _, key := range keys {
		monsterMap, err := redis.StringMap(conn.Do("HGetAll", key))
		if err != nil {
			fmt.Printf("key=%v, err=%v \n", key, err)
			continue
		}
		//将妖怪信息打印在终端
		fmt.Println("======================")
		for k, v := range monsterMap {
			fmt.Printf("%v = %v \n", k, v)
		}
	}
}

//OpenRedisConn 打开Redis链接
func OpenRedisConn() redis.Conn {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return nil
	}
	return conn
}

//练习01：
//1.Monster信息[name,age,skill]
//2.通过终端输入三个monster的信息，使用golang操作redis，存放到redis中(比如使用hash数据类型)
//3.编程，遍历出所有的monster信息，并显示在终端
//4.提示，保存monster可以使用hash数据类型，遍历时先取出所有的keys，比如keys monster*
func main() {
	//1.打开redis链接
	conn := OpenRedisConn()
	if conn == nil {
		return
	}
	defer conn.Close()

	//2.终端输入三个妖怪
	AddMonster(conn)

	//3.遍历获取妖怪信息
	TraversalMonster(conn)

	fmt.Println("exit success")
}
