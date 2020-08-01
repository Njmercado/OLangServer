package handler

import (
	"context"
	"errors"
	"log"

	"github.com/ola/pkg/config/env"
	"github.com/ola/pkg/controller"

	"github.com/ola/pkg/storage/mongodb"
	"github.com/ola/pkg/storage/mongodb/collections"
	"go.mongodb.org/mongo-driver/bson"
)

var mongoDB = mongodb.NewMongoDB("mongodb://localhost:27017", "Ola")
var userCollection = mongoDB.DB.Collection("user")

//GetUser : this function allow me to get some user infor based on its username
func GetUser(username string) (*collections.User, error) {

	user := &collections.User{}

	err := userCollection.FindOne(context.TODO(), bson.M{
		"Username": username,
	}).Decode(&user)

	if err != nil {
		log.Print(err)
		log.Print("This user do not exist")
		return nil, errors.New("This user do not exist")
	}

	return user, nil
}

//UpdateLoggedStatus : allow to update user logged status.
func UpdateLoggedStatus(username string, status bool) error {
	result := userCollection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"Username": username,
		}, bson.M{
			"$set": bson.M{
				"Logged": status,
			},
		}, nil)

	if result.Err() != nil {
		log.Print("Given user couldn't be found")
		return errors.New("Given user couldn't be found")
	} else {
		log.Print("Given user updated correctly")
		return errors.New("Given user updated correctly")
	}
}

func CreateUser(username, password string) error {

	cypherPassword := controller.Encrypt([]byte(password), env.GetPassPhrase())

	user := collections.User{Username: username, Password: string(cypherPassword)}

	result, err := userCollection.InsertOne(
		context.TODO(),
		user,
	)

	if err != nil {
		return err
	}

	log.Print(result.InsertedID)
	return nil
}
