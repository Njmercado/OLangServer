package collections

//User :
type User struct {
	Logged                      bool   `json:"Logged" bson:"Logged"`
	Username                    string `json:"Username" bson:"Username"`
	Password                    string `json:"Password" bson:"Password"`
	Language                    string `json:"Language" bson:"Language"`
	Animal                      `json:"Animal" bson:"Animal"`
	WhereDidYouKnowAboutThisApp string `json:"WhereDidYouKnowAboutThisApp" bson:"WhereDidYouKnowAboutThisApp"`
}
