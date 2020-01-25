package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	user := "root"
	pwd := "123456"
	ip := "127.0.0.1"
	port := "3306"
	db := "test_db"
	dbSource := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		pwd, ip, port, db)
	DB, err := gorm.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
	//设置全局表名禁用复数
	DB.SingularTable(true)
	fmt.Println("db ok!!!")
	return DB
}
