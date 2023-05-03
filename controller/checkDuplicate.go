package controller

import (
	"context"
	"fmt"
	"learning_gin_framework/connection"
	"learning_gin_framework/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckDuplicate(email string) bool {
	var userData model.Users
	setFilter := bson.D{{"email", email}}
	client, err := connection.Connect()
	if err != nil {
		log.Fatal(err)
	}
	userCollection := client.Database("test").Collection("user")
	err = userCollection.FindOne(context.TODO(), setFilter).Decode(&userData)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}
