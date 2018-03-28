// file: web/controllers/user_controller.go

package controllers

import (
	"example.mvc/viewmodels"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

// UserController is our /user controller.
// UserController is responsible to handle the following requests:
// GET  			/user/register
// POST 			/user/register
// GET 				/user/login
// POST 			/user/login
// GET 				/user/profile
// All HTTP Methods /user/logout
type UserController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
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

const userIDKey = "UserID"

var loginStaticView = mvc.View{
	Name: "user/login.html",
	Data: iris.Map{"Title": "User Login"},
}

// GetLogin handles GET: http://localhost:8080/user/login.
func (c *UserController) GetLogin() mvc.Result {

	return loginStaticView
}

// PostLogin handles POST: http://localhost:8080/user/register.
func (c *UserController) PostLogin() mvc.Result {
	var (
		username = c.Ctx.FormValue("UserName")
		password = c.Ctx.FormValue("Password")
	)

	return mvc.Response{
		Path: "/user/profile?username=" + username + "&password=" + password,
	}
}

// GetProfile returns a "Hello {name}" response.
// Demos:
// curl -i http://localhost:8080/user/profile/bee
// curl -i http://localhost:8080/user/profile/anything
func (c *UserController) GetProfile() mvc.Result {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	var user = viewmodels.User{
		Username: username,
		Password: password,
	}
	return mvc.View{
		// if not nil then this error will be shown instead.
		Data: iris.Map{
			"Title": "Profile of " + user.Username,
			"User":  user,
		},
		// redirect to /user/me.
		Name: "user/profile.html",
		// When redirecting from POST to GET request you -should- use this HTTP status code,
		// however there're some (complicated) alternatives if you
		// search online or even the HTTP RFC.
		// Status "See Other" RFC 7231, however iris can automatically fix that
		// but it's good to know you can set a custom code;
		// Code: 303,
	}
}
