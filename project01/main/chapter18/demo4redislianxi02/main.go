package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//OpenRedisConn 打开Redis链接
func OpenRedisConn() redis.Conn {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return nil
	}
	return conn
}

//ShowMenu 显示菜单
func ShowMenu(conn redis.Conn) {
	var key string
	var loop bool = true
	for {
		fmt.Println("-----------------商品浏览历史信息查看-----------------")
		fmt.Println("                 1 添 加 商 品")
		fmt.Println("                 2 查 看 历 史")
		fmt.Println("                 3 退       出")

		fmt.Scanln(&key)
		switch key {
		case "1":
			AddProduct(conn)
		case "2":
			GetTop10Product(conn)
		case "3":
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		if !loop {
			break
		}
	}
}

//AddProduct 添加商品到浏览历史
func AddProduct(conn redis.Conn) {
	var productName string
	fmt.Println("请输入商品名称：")
	fmt.Scanln(&productName)

	_, err := conn.Do("LPush", "productList", productName)
	if err != nil {
		fmt.Println("LPush productList err=", err)
	}
}

//GetTop10Product 显示最近浏览的10个商品名
func GetTop10Product(conn redis.Conn) {
	res, err := redis.Strings(conn.Do("LRange", "productList", 0, 9))
	if err != nil {
		fmt.Println("LRange productList err=", err)
		return
	}

	for index, val := range res {
		fmt.Printf("%d = %v \n", index, val)
	}
}

//练习02:
//1.记录用户浏览商品信息，比如保存商品名
//2.编写一个函数，可以取出某个用户最近浏览的10个商品名
//3.提示：考虑使用list数据类型
func main() {
	//1.打开redis链接
	conn := OpenRedisConn()
	if conn == nil {
		return
	}
	defer conn.Close()

	//2.显示菜单信息
	ShowMenu(conn)

	fmt.Println("exit success")
}
