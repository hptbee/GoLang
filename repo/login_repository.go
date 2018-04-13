package repo

import (
	"fmt"
	"log"

	"example.mvc/context"
	"example.mvc/viewmodels"
	"gopkg.in/mgo.v2/bson"
)

const collection = "users"

// Find a todo by its id
func CheckUsernamePassword(user string, password string) bool {
	todoDB := context.Connect()
	var todo viewmodels.Todo
	err := todoDB.C(collection).Find(bson.M{"userName": user, "password": password}).One(&todo)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func RegisterUser(user viewmodels.User) bool {
	todoDB := context.Connect()
	fmt.Printf("------Insert--------- \n")
	err := todoDB.C(collection).Insert(&user)
	if err != nil {
		return false
	}
	return true
}
