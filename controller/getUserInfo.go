package controller

import (
	"context"
	"learning_gin_framework/model"
	"learning_mongodb/connection"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserInfo(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := connection.Connect()
	if err != nil {
		log.Fatal(err)
	}
	userCollection := client.Database("test").Collection("users")
	filterUser := bson.D{{"Firstname", "Earth"}}
	var userData model.Users
	err = userCollection.FindOne(context.TODO(), filterUser).Decode(&userData)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"data": userData})
	client.Disconnect(ctx)
}
