package controller

import (
	"context"
	"learning_gin_framework/connection"
	"learning_gin_framework/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func HaddleLogin(c *gin.Context) {
	var res model.LoginFilter
	var userData model.Users
	if err := c.BindJSON(&res); err != nil {
		log.Panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	client, err := connection.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userCollection := client.Database("test").Collection("users")
	setFilter := bson.D{{Key: "email", Value: res.Email}}
	err = userCollection.FindOne(context.TODO(), setFilter).Decode(&userData)
	isCompare := CheckPassword(res.Password, userData.Password)

	if err != nil {
		c.JSON(400, gin.H{"data": err})
	} else {
		if isCompare == true {
			token := GenterateToken(res.Email)
			c.JSON(200, gin.H{"token": token})
		} else {
			c.JSON(401, gin.H{"data:": "unauthorized"})
		}
	}
	client.Disconnect(ctx)
}
