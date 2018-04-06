// file: web/controllers/todo_controller.go

package controllers

import (
	"fmt"
	"log"
	"strings"
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

// Find a todo by its id
func SearchByName(name string) (viewmodels.TodoList, error) {
	Connect()
	var todos viewmodels.TodoList
	s := strings.Split(name, "")
	err := todoDB.C(collection).Find(bson.M{"name": bson.M{"$in": s}}).All(&todos)
	return todos, err
}

// Insert a todo into database
func Insert(todo viewmodels.Todo) error {
	Connect()
	fmt.Printf("------Insert--------- \n")
	err := todoDB.C(collection).Insert(&todo)
	return err
}

// Delete By ID an existing todo
func DeleteByID(id bson.ObjectId) error {
	Connect()
	err := todoDB.C(collection).RemoveId(id)
	fmt.Println("delete")
	return err
}

// Update an existing todo
func Update(todo viewmodels.Todo) error {
	Connect()
	fmt.Printf("------Update--------- \n")

	err := todoDB.C(collection).UpdateId(todo.Id, &todo)
	return err
}

// otherwise just return the viewmodels.
var indexStaticView = mvc.View{
	Name: "login/login.html",
	Data: iris.Map{"Title": "User Login"},
}

//Index of Todo
// GET: /todo
func (c *TodoController) Get() mvc.Result {
	result, err := FindAll()
	if err != nil {
		log.Fatal(err)
	}
	return mvc.View{
		Name: "todo/index.html",
		Data: iris.Map{
			"Title": "Todo List",
			"Data":  result,
		},
	}
}

// GetBy returns a todo.
// Demo:
// curl -i http://localhost:8080/todo/1
func (c *TodoController) GetBy(id bson.ObjectId) (todo viewmodels.Todo, found bool) {
	idString := id.String()
	result, err := FindByID(idString)
	if err != nil {
		log.Fatal(err)
	}
	return result, true // it will throw 404 if not found.
}

// PostTodo handles POST: http://localhost:8080/todo/insert.
func (c *TodoController) PostInsert() mvc.Result {

	name := c.Ctx.PostValue("name")
	completed := c.Ctx.PostValue("completed")
	bcompleted := false
	if completed == "on" {
		bcompleted = true
	}
	fmt.Printf(name)
	// completed := c.Ctx.FormValue("completed")
	// fmt.Printf(completed)
	// // get the request data for poster and genre
	// boolCom, err := strconv.ParseBool(completed)
	// if err != nil {
	// 	// handle the error in some way
	// }
	model := viewmodels.Todo{
		Id:        bson.NewObjectId(),
		Name:      name,
		Completed: bcompleted,
	}
	err2 := Insert(model)
	if err2 != nil {
		log.Fatal(err2)
	}
	return mvc.Response{
		Path: "/todo",
	}
}

// PostTodo handles POST: http://localhost:8080/todo/insert.
func (c *TodoController) PostUpdate() mvc.Result {

	id := c.Ctx.PostValue("id")
	name := c.Ctx.PostValue("name")
	completed := c.Ctx.PostValue("completed")
	bcompleted := false
	if completed == "on" {
		bcompleted = true
	}

	fmt.Println(completed)
	fmt.Println(name)
	// completed := c.Ctx.FormValue("completed")
	// fmt.Printf(completed)
	// // get the request data for poster and genre
	// boolCom, err := strconv.ParseBool(completed)
	// if err != nil {
	// 	// handle the error in some way
	// }
	model := viewmodels.Todo{
		Id:        bson.ObjectIdHex(id),
		Name:      name,
		Completed: bcompleted,
	}
	err2 := Update(model)
	if err2 != nil {
		log.Fatal(err2)
	}
	return mvc.Response{
		Path: "/todo",
	}
}

// GetRemove deletes a todo.
// Demo:
// GET: /todo/remove/{id:string}
func (c *TodoController) GetRemoveBy(id string) mvc.Result {
	fmt.Println(id)
	var idA = bson.ObjectIdHex(id)
	fmt.Println("convert Id")
	wasDel := DeleteByID(idA)

	if wasDel != nil {
		// return the deleted todo's ID
		log.Fatal(wasDel)
	}
	return mvc.Response{
		Path: "/todo",
	}
}

//Post search
func (c *TodoController) PostSearch() mvc.Result {
	keyword := c.Ctx.PostValue("keyword")

	result, err := SearchByName(keyword)
	if err != nil {
		log.Fatal(err)
	}
	return mvc.View{
		Name: "todo/index.html",
		Data: iris.Map{
			"Title": "Todo List",
			"Data":  result,
		},
	}
}
