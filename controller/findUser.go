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

func FindingUser(c *gin.Context) {
	var res model.UserFilter
	var userData model.Users

	if err := c.BindJSON(&res); err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	client, err := connection.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userCollection := client.Database("test").Collection("users")
	setFilter := bson.D{{"Firstname", res.Firstname}}
	err = userCollection.FindOne(context.TODO(), setFilter).Decode(&userData)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"data": userData})
	client.Disconnect(ctx)
}
