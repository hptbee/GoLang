package repo

import (
	"log"
	"time"

	"example.mvc/viewmodels"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var todoDB *mgo.Database

const (
	hosts      = "ds153948.mlab.com:53948"
	database   = "todos_db"
	username   = "trungtp"
	password   = "trungtp123"
	collection = "todos"
)

type TodoRepository struct {
	// Our TodoService, it's an interface which
	// is binded from the main application.
}

// Query represents the visitor and action queries.
type Query func(viewmodels.Todo) bool

// Conect To Mongo DB
func Connect() {
	infos := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(infos)
	if err != nil {
		log.Fatal(err)
	}

	todoDB = session.DB(database)
}

// Find list of todos
func FindAll() (viewmodels.TodoList, error) {
	Connect()
	var todos viewmodels.TodoList
	err := todoDB.C(collection).Find(bson.M{}).All(&todos)
	return todos, err
}

// Find a todo by its id
func FindByID(id string) (viewmodels.Todo, error) {
	Connect()
	var todo viewmodels.Todo
	err := todoDB.C(collection).FindId(bson.ObjectIdHex(id)).One(&todo)
	return todo, err
}

// Insert a todo into database
func Insert(todo viewmodels.Todo) error {
	Connect()
	err := todoDB.C(collection).Insert(&todo)
	return err
}

// Delete an existing todo
func Delete(todo viewmodels.Todo) error {
	Connect()
	err := todoDB.C(collection).Remove(&todo)
	return err
}

// Update an existing todo
func Update(todo viewmodels.Todo) error {
	Connect()
	err := todoDB.C(collection).UpdateId(todo.Id, &todo)
	return err
}
