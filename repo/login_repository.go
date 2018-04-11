package repo

import (
	"fmt"

	"example.mvc/context"
	"example.mvc/viewmodels"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "users"

var todoDB *mgo.Database

// Find a todo by its id
func CheckUsernamePassword(user string, password string) bool {
	todoDB = context.Connect()
	var todo viewmodels.Todo
	err := todoDB.C(collection).Find(bson.M{"userName": user, "password": password}).One(&todo)
	if err != nil {
		return false
	}
	return true
}

func RegisterUser(user viewmodels.User) bool {
	todoDB = context.Connect()
	fmt.Printf("------Insert--------- \n")
	err := todoDB.C(collection).Insert(&user)
	if err != nil {
		return false
	}
	return true
}