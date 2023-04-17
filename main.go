package main

import (
	"fmt"
	"learning_gin_framework/controller"

	"github.com/gin-gonic/gin"
)

const port string = ":9888"

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/member", func(c *gin.Context) {
		data := controller.AddingMember()
		c.JSON(200, data)
	})

	router.GET("/getUser", controller.GetUserInfo)
	router.POST("/add/member", controller.TestPostMethod)

	
	err := router.Run(port)
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
	fmt.Println("Framework gin server start at http://localhost", port)
}
