package router

import (
	"gin-app/control"
	"github.com/gin-gonic/gin"
)

func InitRouter() {

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sss",
		})
	})

	//分组路由
	v1 := r.Group("v1/user")
	{
		v1.GET("/", control.FetchAll)
		v1.GET("/:id", control.FetchSingleUser)
	}
	r.Run()
}
