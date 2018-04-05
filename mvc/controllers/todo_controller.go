// file: web/controllers/todo_controller.go

package controllers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"example.mvc/viewmodels"
)

var todoDB *mgo.Database

const (
	hosts      = "ds153948.mlab.com:53948"
	database   = "todos_db"
	username   = "trungtp"
	password   = "trungtp123"
	collection = "todos"
)

// TodoController is our /todos controller.
type TodoController struct {
	Ctx iris.Context
}

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

// otherwise just return the viewmodels.
var indexStaticView = mvc.View{
	Name: "login/login.html",
	Data: iris.Map{"Title": "User Login"},
}

func (c *TodoController) Get() mvc.Result {
	fmt.Printf("------controller--------- \n")
	result, err := FindAll()
	if err != nil {
		log.Fatal(err)
	}
	return mvc.View{
		Name: "todo/index.html",
		Data: result,
	}
}

// GetBy returns a todo.
// Demo:
// curl -i http://localhost:8080/todos/1
func (c *TodoController) GetBy(id bson.ObjectId) (todo viewmodels.Todo, found bool) {
	idString := id.String()
	result, err := FindByID(idString)
	if err != nil {
		log.Fatal(err)
	}
	return result, true // it will throw 404 if not found.
}

// PutBy updates a todo.
// Demo:
// curl -i -X PUT -F "genre=Thriller" -F "poster=@/Users/kataras/Downloads/out.gif" http://localhost:8080/todos/1
func (c *TodoController) Insert() mvc.Result {
	var (
		name      = c.Ctx.FormValue("name")
		completed = c.Ctx.FormValue("completed")
	)
	// get the request data for poster and genre
	boolCom, err := strconv.ParseBool(completed)
	if err != nil {
		// handle the error in some way
	}
	model := viewmodels.Todo{
		Id:        bson.NewObjectId(),
		Name:      name,
		Completed: boolCom,
	}
	err2 := Update(model)
	if err2 != nil {
		log.Fatal(err2)
	}
	return mvc.Response{
		Path: "/todos",
	}
}

// DeleteBy deletes a todo.
// Demo:
// curl -i -X DELETE -u admin:password http://localhost:8080/todos/1
// func (c *TodoController) DeleteBy(id bson.ObjectId) interface{} {
// 	wasDel := DeleteByID(id)
// 	if wasDel {
// 		// return the deleted todo's ID
// 		return iris.Map{"deleted": id}
// 	}
// 	// right here we can see that a method function can return any of those two types(map or int),
// 	// we don't have to specify the return type to a specific type.
// 	return iris.StatusBadRequest
// }
