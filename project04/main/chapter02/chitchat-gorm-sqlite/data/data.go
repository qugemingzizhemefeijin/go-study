package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var Db *gorm.DB
var SqlDB *sql.DB

// https://github.com/boyiling123/gohub/tree/1cf0ee6755df359cdf68e2ec90148751b0ea567f
// git@github.com:boyiling123/gohub.git gorm 连接池使用案例

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
	Db ,err = gorm.Open(sqlite.Open("E:/chitchat.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// 开启调式
	Db = Db.Debug()

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	SqlDB, err = Db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	SqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	SqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	SqlDB.SetConnMaxLifetime(time.Hour)
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
