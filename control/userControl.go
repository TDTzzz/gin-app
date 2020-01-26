package control

import (
	"fmt"
	"gin-app/database"
	"gin-app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FetchSingleUser(c *gin.Context) {
	var user model.User
	id := c.Param("id")
	fmt.Println(id)
	dbConn := database.InitDb()
	dbConn.First(&user, id)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "user": user})
}

func FetchAll(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
