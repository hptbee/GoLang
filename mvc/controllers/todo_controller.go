// file: web/controllers/todo_controller.go

package controllers

import (
	"fmt"
	"log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/mgo.v2/bson"

	"example.mvc/repo"
	"example.mvc/viewmodels"
)

type TodoController struct {
	Ctx iris.Context
}

// otherwise just return the viewmodels.
var indexStaticView = mvc.View{
	Name: "login/login.html",
	Data: iris.Map{"Title": "User Login"},
}

//Index of Todo
// GET: /todo
func (c *TodoController) Get() mvc.Result {
	result, err := repo.FindAll()
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
	result, err := repo.FindByID(idString)
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
	err2 := repo.Insert(model)
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
	err2 := repo.Update(model)
	if err2 != nil {
		log.Fatal(err2)
	}
	return mvc.Response{
		Path: "/todo",
	}
}

// GetRemoveBy deletes a todo.
// Demo:
// GET: /todo/remove/{id:string}
func (c *TodoController) GetRemoveBy(id string) mvc.Result {
	idA := bson.ObjectIdHex(id)
	wasDel := repo.DeleteByID(idA)

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

	result, err := repo.SearchByName(keyword)
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
