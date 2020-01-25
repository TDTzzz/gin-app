package router

import (
	"fmt"
	"gin-app/database"
	"gin-app/model"
	"github.com/gin-gonic/gin"
)

func InitRouter() {

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sss",
		})
	})

	r.GET("/db", func(c *gin.Context) {
		var user model.User
		dbConn := database.InitDb()
		dbConn.First(&user)
		fmt.Println(user)
		c.JSON(200, gin.H{
			"user": user,
		})
	})
	r.Run()
}
