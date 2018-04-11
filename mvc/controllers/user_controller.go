// file: web/controllers/user_controller.go

package controllers

import (
	"example.mvc/authen"
	"example.mvc/viewmodels"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// UserController is our /user controller.
// UserController is responsible to handle the following requests:
// GET 				/user/profile
// All HTTP Methods /user/logout
type UserController struct {
	Ctx iris.Context
}

const userIDKey = "UserID"

//Get Current User Id from Session

// GetProfile returns a "Hello {name}" response.
// Demos:
// curl -i http://localhost:8080/user/profile/bee
// curl -i http://localhost:8080/user/profile/anything
func (c *UserController) GetProfile() mvc.Result {

	var user = viewmodels.User{
		Username: "username",
		Password: "password",
		Name:     "Tùng",
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

func (c *UserController) GetLogout() mvc.Result {
	authen.RemoveSession("userid", c.Ctx)
	return mvc.Response{
		Path: "/login",
	}
}
