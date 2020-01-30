package main

import (
	"fmt"
	"gin-app/database"
	"gin-app/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dbConn := database.InitDb()
	//fmt.Println(dbConn.HasTable(&model.HouseInfo{}))
	res := dbConn.CreateTable(&model.HouseInfo{})
	fmt.Println(res)
	//router.InitRouter()
}
