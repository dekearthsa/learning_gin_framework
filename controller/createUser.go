package controller

import (
	"context"
	"fmt"
	"learning_gin_framework/connection"
	"learning_gin_framework/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var res model.Users
	if err := c.BindJSON(&res); err != nil {
		log.Fatal(err)
	}

	duplicateUser := CheckDuplicate(res.Email)
	fmt.Println("duplicateUser => ", duplicateUser)
	if duplicateUser == true {
		c.JSON(400, gin.H{"status": "This user alreadly register."})
	} else {
		hash, err := HashPassword(res.Password)
		if err != nil {
			c.JSON(400, gin.H{"status": "Bad Request"})
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		client, err := connection.Connect()
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"status": "Internal Server Error"})
		}
		userCollection := client.Database("test").Collection("users")
		userCreator := model.Users{Firstname: res.Firstname, Lastname: res.Lastname, Email: res.Email, Password: hash}
		insertResult, err := userCollection.InsertOne(context.TODO(), userCreator)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("insert success. ", insertResult)
			c.JSON(200, gin.H{"status": "OK"})
		}
		client.Disconnect(ctx)
	}

}
