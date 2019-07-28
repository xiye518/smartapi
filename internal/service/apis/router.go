package apis

import (
	"github.com/gin-gonic/gin"
)

func RunGinService() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,Cookie")
	})

	router.Use(CheckAuth())

	router.GET("/user", GetUser)

	router.POST("/user", AddUser)

	router.PUT("/user/:id", UpdateUser)

	router.DELETE("/user/:id", DeleteUser)

	router.GET("/user/list", GetUserList)

	return router
}
