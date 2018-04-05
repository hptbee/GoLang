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

//Get Current User Id from Session
func (c *UserController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

// GetProfile returns a "Hello {name}" response.
// Demos:
// curl -i http://localhost:8080/user/profile/bee
// curl -i http://localhost:8080/user/profile/anything
func (c *UserController) GetProfile() mvc.Result {
	userId := c.getCurrentUserID()
	if userId == 0 {
		return mvc.Response{
			Path: "/login",
		}
	}

	var user = viewmodels.User{
		Username:  "username",
		Password:  "password",
		Firstname: "Tùng",
	}
	return mvc.View{
		// if not nil then this error will be shown instead.
		Data: iris.Map{
			"Title": "Profile of " + user.Username,
			"User":  user,
		},
		// redirect to /user/me.
		Name: "user/profile.html",
		
	}
}
