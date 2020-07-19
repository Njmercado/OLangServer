package collections

import "go.mongodb.org/mongo-driver/bson/primitive"

//User :
type User struct {
	ID                          primitive.ObjectID `json:"_id" bson:"_id"`
	Logged                      bool               `json:"logged" bson:"logged"`
	Username                    string             `json:"username" bson:"username"`
	Password                    string             `json:"password" bson:"password"`
	language                    string
	Animal                      `json:"animal" bson:"animal"`
	WhereDidYouKnowAboutThisApp string `json:"whereDidYouKnowAboutThisApp" bson:"whereDidYouKnowAboutThisApp"`
}
