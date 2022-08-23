package main

import (
	"database/sql"
	// 加 _ 就会初始化包里 init() 函数。
	_ "github.com/go-sql-driver/mysql"
	// 添加 . 这算是别名的一个特殊方式吧，可以省略包名直接调用该包的方法，
	// 例如 printf 一般都是 fmt.printf("hello") 来使用，但是在引用之前加了 . ，就可以去掉 fmt ，而直接使用 printf 。
	"log"
	// 直接在引用的前面加一个字符或者字符串，在使用包里面的函数时就可以直接使用别名再点出函数名。
)

// Go 官方提供了 database/sql 包来给用户进行和数据库打交道的工作，database/sql 库实际只提供了一套操作数据库的接口和规范，
// 例如抽象好的 SQL 预处理（prepare），连接池管理，数据绑定，事务，错误处理等等。官方并没有提供具体某种数据库实现的协议支持。

/**
例如 MySQL 打交道，还需要再引入 MySQL 的驱动，像下面这样：

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")
import _ "github.com/go-sql-driver/mysql"
这条 import 语句会调用了 mysql 包的 init 函数，做的事情也很简单：

func init() {
	sql.Register("mysql", &MySQLDriver{})
}
 */

func main() {
	// db 是一个sql.DB类型的对象
	// 该对象线程安全，且内部已包含了一个连接池
	// 连接池的选项可以在 sql.DB的方法中设置，这里为了简单省略了。

	db, err := sql.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var (
		id int
		name string
	)
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 必须要把rows里的内容读完，或者显示调用Close()方法，
	// 否则在 defer 的 rows.Close() 执行之前，连接永远不会释放
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}