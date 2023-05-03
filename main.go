package main

import (
	"fmt"
	"learning_gin_framework/controller"

	"github.com/gin-gonic/gin"
)

const port string = ":9888"

func main() {

	router := gin.New()
	protected := router.Group("/")

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

	// ทดลอง POST ไป ตาม username ต่อกับ database เพื่อหาชื่อ user //
	router.POST("/findUser", controller.FindingUser)
	// ทดลอง POST ไป ต่อกับ database เพื่อสร้าง user //
	router.POST("/createUser", controller.CreateUser)
	router.POST("/add/member", controller.TestPostMethod)
	router.POST("/login")

	err := router.Run(port)
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
	fmt.Println("Framework gin server start at http://localhost", port)
}
