package handler

import (
	"errors"
	"time"

	"github.com/ola/pkg/config/env"
	"github.com/ola/pkg/controller"
)

// var mongoDB = mongodb.NewMongoDB("mongodb://localhost:27017", "Ola")
// var userCollection = mongoDB.DB.Collection("cookie")

var listOfCookies = make(map[string]Cookie)

//Cookie creates a cookie structure to manage all its information on header.
//CookieName is refers to the given name which gonna be spotted as cookie.
//LimitTime refers to life time than this cookie gonna be alive. This field are represented by seconds.
type Cookie struct {
	CookieName string
	LimitTime  int64
}

// General functions
func thisCookieExist(cookieID string) bool {
	return len(listOfCookies[cookieID].CookieName) > 0
}
func thisCookieHasExceededTheTime(cookieID string) bool {
	return listOfCookies[cookieID].LimitTime >= time.Now().Unix()
}

//GenerateID create a new ID, for this cookie, based on global key and username
//This is a unique value than if it is changed it gonna be rejected
func GenerateID(username string) string {
	//TODO: improve this zone, I'm do not secure of this function.
	//Create a better and secure, cookie key, generator.
	globalKey := env.GetPassPhrase()
	return controller.Hash(globalKey + username)
}

//NewCookie create a new key based on name given from any zone in this project
func NewCookie(name string, limitTime int) (Cookie, error) {

	id := GenerateID(name)

	if thisCookieExist(id) {
		return Cookie{}, errors.New("This cookie already exist")
	}

	limitTimePlusCurrentTime := int64(time.Duration(limitTime)*time.Second) + time.Now().Unix()
	//Create cookie
	cookie := Cookie{name, limitTimePlusCurrentTime}

	//Add new row into cookie's list
	listOfCookies[GenerateID(name)] = cookie

	//Watch cookie lifetime
	go watchCookie(name)

	return cookie, nil
}

//GetCookie allows me to retrieve some specific cookie.
func GetCookie(name string) (Cookie, error) {
	id := GenerateID(name)
	if thisCookieExist(id) {
		return listOfCookies[id], nil
	}
	return Cookie{}, errors.New("This cookie does not exist")
}

//DeleteCookie allows me to delete some cookie based on its name
func DeleteCookie(name string) error {
	id := GenerateID(name)
	if thisCookieExist(id) {
		delete(listOfCookies, id)
		return nil
	}
	return errors.New("This cookie does not exist")
}

//This function allows me to manage cookie lifetime through goroutines.
//Once time has been reached it delete given cookie
func watchCookie(name string) {
	id := GenerateID(name)
	cookie := listOfCookies[id]
	cookieLimitTime := time.Duration(cookie.LimitTime) * time.Second
	time.AfterFunc(cookieLimitTime, func() { DeleteCookie(name) })
}
