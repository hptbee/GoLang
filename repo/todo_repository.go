package repo

import (
	"fmt"
	"strings"

	"example.mvc/context"
	"example.mvc/viewmodels"
	"gopkg.in/mgo.v2/bson"
)

const collectionTodo = "todos"

func FindAll() (viewmodels.TodoList, error) {
	todoDB := context.Connect()
	var todos viewmodels.TodoList
	err := todoDB.C(collection).Find(bson.M{}).All(&todos)
	return todos, err
}

// Find a todo by its id
func FindByID(id string) (viewmodels.Todo, error) {
	todoDB := context.Connect()
	var todo viewmodels.Todo
	err := todoDB.C(collectionTodo).FindId(bson.ObjectIdHex(id)).One(&todo)
	return todo, err
}

// Find a todo by its id
func SearchByName(name string) (viewmodels.TodoList, error) {
	todoDB := context.Connect()
	var todos viewmodels.TodoList
	s := strings.Split(name, "")
	err := todoDB.C(collectionTodo).Find(bson.M{"name": bson.M{"$in": s}}).All(&todos)
	return todos, err
}

// Insert a todo into database
func Insert(todo viewmodels.Todo) error {
	todoDB := context.Connect()
	fmt.Printf("------Insert--------- \n")
	err := todoDB.C(collectionTodo).Insert(&todo)
	return err
}

// Delete By ID an existing todo
func DeleteByID(id bson.ObjectId) error {
	todoDB := context.Connect()
	err := todoDB.C(collectionTodo).RemoveId(id)
	fmt.Println("delete")
	return err
}

// Update an existing todo
func Update(todo viewmodels.Todo) error {
	todoDB := context.Connect()
	fmt.Printf("------Update--------- \n")

	err := todoDB.C(collectionTodo).UpdateId(todo.Id, &todo)
	return err
}
