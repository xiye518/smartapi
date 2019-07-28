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

	v1 := router.Group("/v1/api")
	v1.Use(CheckAuth())

	v1.GET("/user", GetUser)
	v1.POST("/user", AddUser)
	v1.PUT("/user/:id", UpdateUser)
	v1.DELETE("/user/:id", DeleteUser)
	v1.GET("/user/list", GetUserList)

	return router
}
