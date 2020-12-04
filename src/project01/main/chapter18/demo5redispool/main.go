package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接的代码
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//先从连接池中取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "透明母猫!")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	//取出
	res, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	fmt.Println("Name = ", res)

	//如果我们要从pool取出链接，一定要保证连接池是没有关闭的，否则在DO的时候，会报错
	pool.Close()

	//conn2 := pool.Get()
	//fmt.Println("conn2=", conn2)

	fmt.Println("exit success")
}
