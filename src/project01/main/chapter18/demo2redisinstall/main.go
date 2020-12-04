package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis" //引入redis包
)

//如果安装Go Redis库
func main() {
	//1. 使用第三方开源的redis库：https://github.com/gomodule/redigo
	//2. 在使用Redis前，先安装第三方Redis库，在GOPATH路径下执行安装指令：
	//   go get github.com/gomodule/redigo
	// 安装成功后，就可以使用了

	//1.这里测试一下链接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭...

	fmt.Println("conn success ", conn)

	//=============================操作kv=====================================================

	//2. 通过go 向redis中写入数据 string [key-val]
	_, err = conn.Do("Set", "name", "tomcat")
	if err != nil {
		fmt.Println("Set err=", err)
		return
	}

	//3. 通过go 向redis读取数据 string [key-val]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("Get err=", err)
		return
	}

	//因为返回r是interface{}
	//因为 name 对应的值是stirng，因此我们需要转换
	//nameString := r.(string) //直接转会报错： interface conversion: interface {} is []uint8, not string
	//需要通过 redis.String() 函数来执行
	fmt.Println("res=", r)

	//=============================操作hash===================================================
	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("HSet Name err=", err)
		return
	}
	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("HSet Age err=", err)
		return
	}
	_, err = conn.Do("HMSet", "user01", "score", 99.9, "address", "北京市朝阳区")
	if err != nil {
		fmt.Println("HMSet err=", err)
		return
	}

	//获取hash信息
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("HGet Name err=", err)
		return
	}
	r2, err := redis.String(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("HGet Age err=", err)
		return
	}

	fmt.Printf("HGet Name=%v, Age=%v \n", r1, r2)

	r3, err := redis.Strings(conn.Do("HGetAll", "user01")) //返回的是切片
	if err != nil {
		fmt.Println("HMGet user01 err=", err)
		return
	}
	fmt.Println("===================================")
	for i := 0; i < len(r3); i = i + 2 {
		fmt.Printf("%v = %v \n", r3[i], r3[i+1])
	}
	fmt.Println("===================================")

	fmt.Println("exit success")
}
