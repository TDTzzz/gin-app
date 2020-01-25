package main

import (
	"gin-app/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//dbConn := database.InitDb()
	//var loc *time.Location

	//dbConn.CreateTable(&model.User{})
	//dbConn.Create(&model.User{
	//	Name:     "wcy",
	//	Age:      sql.NullInt64{Int64: int64(23)},
	//	Birthday: birthday,
	//})
	//var user model.User
	//dbConn.First(&user,2)
	//fmt.Println(user.Name)

	router.InitRouter()
}
