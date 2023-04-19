package controller

import (
	"context"
	"fmt"
	"learning_gin_framework/model"
	"learning_mongodb/connection"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var res model.Users
	if err := c.BindJSON(&res); err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	client, err := connection.Connect()
	if err != nil {
		log.Fatal(err)
	}
	userCollection := client.Database("test").Collection("users")
	userCreator := model.Users{res.Firstname, res.Lastname, res.Email, res.Password}
	insertResult, err := userCollection.InsertOne(context.TODO(), userCreator)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insert success. ", insertResult)
	client.Disconnect(ctx)
}
