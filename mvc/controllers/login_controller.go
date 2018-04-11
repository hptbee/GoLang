// file: web/controllers/login_controller.go
package controllers

import (
	"example.mvc/authen"
	context "example.mvc/repo"
	"example.mvc/viewmodels"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"gopkg.in/mgo.v2/bson"
)

// LoginController is our /login controller.
// LoginController is responsible to handle the following requests:
// GET 				/login
// POST 			/login
type LoginController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new LoginController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	//Service services.UserService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

// GetLogin handles GET: http://localhost:8080/login.
func (c *LoginController) Get() mvc.Result {
	return mvc.View{
		Name: "login/login.html",
		Data: iris.Map{"Title": "User Login"},
	}
}

// PostLogin handles POST: http://localhost:8080/login.
func (c *LoginController) Post() mvc.Result {
	var (
		username = c.Ctx.FormValue("UserName")
		password = c.Ctx.FormValue("Password")
	)

	if context.CheckUsernamePassword(username, password) {
		authen.CreateSession("userid", "23", c.Ctx)

		return mvc.Response{
			Path: "/user/profile",
		}
	}

	return mvc.Response{
		Path: "/login",
	}
}

func (c *LoginController) GetRegister() mvc.Result {
	return mvc.View{
		Name: "login/register.html",
		Data: iris.Map{"Title": "User Login"},
	}
}

func (c *LoginController) PostRegister() mvc.Result {
	var user = viewmodels.User{
		Id:       bson.NewObjectId(),
		Username: c.Ctx.FormValue("UserName"),
		Password: c.Ctx.FormValue("Password"),
		Name:     c.Ctx.FormValue("Name"),
		Email:    c.Ctx.FormValue("Email"),
	}
	if context.RegisterUser(user) {
		authen.CreateSession("userid", "user.Id.Hex()", c.Ctx)
		return mvc.Response{
			Path: "/user/profile",
		}
	}
	return mvc.View{
		Name: "login/register.html",
		Data: iris.Map{"Title": "User Login"},
	}
}
