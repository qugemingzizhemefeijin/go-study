package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
CREATE TABLE `food` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID，商品Id',
  `name` varchar(30) NOT NULL COMMENT '商品名',
  `price` decimal(10,2) unsigned  NOT NULL COMMENT '商品价格',
  `type_id` int(10) unsigned NOT NULL COMMENT '商品类型Id',
  `create_time` int(10) NOT NULL DEFAULT 0 COMMENT '创建时间',
   PRIMARY KEY (`id`)
  ) ENGINE=INNODB DEFAULT CHARSET=utf8;
 */

// 默认gorm对struct字段名使用Snake Case命名风格转换成mysql表字段名(需要转换成小写字母)。
// 根据gorm的默认约定，上面例子只需要使用gorm:"column:create_time"标签定义为CreateTime字段指定表字段名，其他使用默认值即可。
// 提示：Snake Case命名风格，就是各个单词之间用下划线（_）分隔，例如： CreateTime的Snake Case风格命名为create_time

//字段注释说明了gorm库把struct字段转换为表字段名长什么样子。
type Food struct {
	Id         int  //表字段名为：id
	Name       string //表字段名为：name
	Price      float64 //表字段名为：price
	TypeId     int  //表字段名为：type_id

	//字段定义后面使用两个反引号``包裹起来的字符串部分叫做标签定义，这个是golang的基础语法，不同的库会定义不同的标签，有不同的含义
	CreateTime int64 `gorm:"column:create_time"`  //表字段名为：create_time
}

//设置表名，可以通过给Food struct类型定义 TableName函数，返回一个字符串作为表名
// 建议: 默认情况下都给模型定义表名，有时候定义模型只是单纯的用于接收手写sql查询的结果，这个时候是不需要定义表名；
// 手动通过gorm函数Table()指定表名，也不需要给模型定义TableName函数。
func (v Food) TableName() string {
	return "food"
}

/**
// GORM 定义一个 gorm.Model 结构体，其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt。
// gorm.Model 的定义
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 以将它嵌入到我们的结构体中，就以包含这几个字段，类似继承的效果。
type User struct {
  gorm.Model // 嵌入gorm.Model的字段
  Name string
}

GORM 约定使用 CreatedAt、UpdatedAt 追踪创建/更新时间。如果定义了这种字段，GORM 在创建、更新时会自动填充当前时间。
要使用不同名称的字段，您可以配置 autoCreateTime、autoUpdateTime 标签
如果想要保存 UNIX（毫/纳）秒时间戳，而不是 time，只需简单地将 time.Time 修改为 int 即可。

type User struct {
  CreatedAt time.Time // 默认创建时间字段， 在创建时，如果该字段值为零值，则使用当前时间填充
  UpdatedAt int       // 默认更新时间字段， 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
  Updated   int64 `gorm:"autoUpdateTime:nano"` // 自定义字段， 使用时间戳填纳秒数充更新时间
  Updated   int64 `gorm:"autoUpdateTime:milli"` //自定义字段， 使用时间戳毫秒数填充更新时间
  Created   int64 `gorm:"autoCreateTime"`      //自定义字段， 使用时间戳秒数填充创建时间
}
 */

func print(food Food) {
	data, err := json.Marshal(food)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	//输出序列化后的结果
	fmt.Printf("序列化后=%v\n", string(data))
}

//定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

//包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init() {
	//配置MySQL连接参数
	username := "root"  //账号
	password := "abcd1234" //密码
	host := "127.0.0.1" //数据库地址，可以是Ip或者域名
	port := 3306 //数据库端口
	Dbname := "test" //数据库名
	timeout := "10s" //连接超时，10秒

	// 拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，
	// 我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)


	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	// 全局变量*gorm.DB,只是一个对数据库连接前的初始化加载,你不需要关心数据库连接关闭的问题。 但是你需要关注数据库连接数的健康状况
	// 连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	// 非连接池
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("连接数据库失败, error=" + err.Error())
	//}

	// 初始化连接池
	sqlDB, _ := _db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(10)   //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(2)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

// 获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
// 不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return _db
}

func main() {
	cake := Food{}
	// var cake Food
	GetDB().Debug().Where("name = ?", "admin").First(&cake)

	print(cake)
}