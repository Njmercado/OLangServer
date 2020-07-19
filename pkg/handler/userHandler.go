package handler

import (
	"context"
	"log"

	"github.com/ola/pkg/storage/mongodb"
	"github.com/ola/pkg/storage/mongodb/collections"
	"go.mongodb.org/mongo-driver/bson"
)

var mongoDB = mongodb.NewMongoDB("mongodb://localhost:27017", "Ola")
var userCollection = mongoDB.DB.Collection("user")

//GetUser : this function allow me to get some user infor based on its username
func GetUser(username string) *collections.User {

	user := &collections.User{}

	err := userCollection.FindOne(context.TODO(), bson.M{
		"username": username,
	}).Decode(&user)

	if err != nil {
		log.Print(err)
		log.Print("This user do not exist")
		return nil
	}

	return user
}

//UpdateLoggedStatus : allow to update user logged status.
func UpdateLoggedStatus(username string, status bool) {
	result := userCollection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"username": username,
		}, bson.M{
			"$set": bson.M{
				"logged": status,
			},
		}, nil)

	if result.Err() != nil {
		log.Print("Given user couldn't be found")
	} else {
		log.Print("Given user updated correctly")
	}
}
