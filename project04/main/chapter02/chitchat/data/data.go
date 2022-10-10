package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

// init()函数是一个特殊的函数，存在以下特性：
// 1.不能被其他函数调用，而是在main函数执行之前，自动被调用；
// 2.init函数不能作为参数传入；
// 3.不能有传入参数和返回值；

// 变量除了可以在全局声明中初始化，也可以在 init ()函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，
// 而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。
// 每个源文件可以包含多个init函数，而且会按先后顺序执行，优先级高于main。初始化总是以单线程执行，并且按照包的依赖关系顺序执行。
// 一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性。
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	Db.SetMaxOpenConns(3)//连接池最大连接数
	Db.SetMaxIdleConns(2)//空闲时候保持的连接数
	return
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
