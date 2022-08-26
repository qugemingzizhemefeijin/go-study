package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/**
CREATE TABLE Product(
   ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
   CreatedAt DATETIME NOT NULL,
   UpdatedAt DATETIME NOT NULL,
   DeletedAt DATETIME NOT NULL,
   Code      VARCHAR(50)    NOT NULL,
   Price     INT     NOT NULL
);

insert into Product(CreatedAt,UpdatedAt,DeletedAt,Code,Price)
values(CURRENT_TIMESTAMP,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP, 'A01', 100);
 */

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func print(product Product) {
	data, err := json.Marshal(product)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	//输出序列化后的结果
	fmt.Printf("序列化后=%v\n", string(data))
}

func main() {
	db ,err := gorm.Open(sqlite.Open("E:/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect datebase")
	}

	db = db.Debug()

	// 自动创建表
	// 如果使用此来自动创建表，注意，结构体的字段必须要大写。
	/**
	type User struct {
	  gorm.Model
	  Name         string
	  Age          sql.NullInt64
	  Birthday     *time.Time
	  Email        string  `gorm:"type:varchar(100);unique_index"`
	  Role         string  `gorm:"size:255"` // 设置字段大小为255
	  MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	  Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	  Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	  IgnoreMe     int     `gorm:"-"` // 忽略本字段
	}
	 */
	_ = db.AutoMigrate(&Product{})

	// 创建记录
	db.Create(&Product{
		Model: gorm.Model{
			ID: uint(2),
		},
		Code: "D42",
		Price: 101,
	})

	// 读取
	var product Product
	db.First(&product, 2) // 通过主键查询
	db.First(&product, "code = ?", "D42") // 通过属性来查询
	print(product)

	product = Product{}
	db.Take(&product) // 不排序查询第一条数据
	print(product)

	product = Product{}
	result := db.Last(&product) // 根据主键排序返回最后一条数据

	println(result.RowsAffected)  // returns count of records found
	println(result.Error)         // returns error or nil)

	// check error ErrRecordNotFound
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	// 如果你想避免 ErrRecordNotFound 错误，你可以使用 Find like db.Limit(1).Find(&user)，Find 方法接受结构和切片数据
	db.Limit(1).Find(&product)

	ret := map[string]interface{}{}
	db.Model(&Product{}).First(&ret)

	// doesn't work
	//result := map[string]interface{}{}
	// db.Table("users").First(&result)

	ret = map[string]interface{}{}
	db.Table("products").Take(&ret)

	// no primary key defined, results will be ordered by first field (i.e., `Code`)
	type Language struct {
		Code string
		Price uint
	}
	db.Table("products").First(&Language{})
	// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1

	productList := make([]Product,0)
	db.Find(&productList, []int{1,2,3})
	data, _ := json.Marshal(productList)
	//输出序列化后的结果
	fmt.Printf("序列化列表=%v\n", string(data))

	db.Where("code = 'aa'").Or(Product{Code: "bb", Price: 200}).Find(&productList)
	data, _ = json.Marshal(productList)
	//输出序列化后的结果
	fmt.Printf("序列化列表=%v\n", string(data))

	// 修改
	db.Model(&product).Update("Price", 200) // 修改价格
	print(product)
	// 一次修改多个属性
	db.Model(&product).Updates(Product{Price: 201, Code: "F42"}) // 必须要设置字段的，否则应该会报错
	print(product)
	db.Model(&product).Updates(map[string]interface{}{"Price": 300, "Code": "H48"})
	print(product)

	// 逻辑删除数据
	// db.Delete(&product, 2)
	db.Unscoped().Delete(&product, 2)
}