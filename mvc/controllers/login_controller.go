// file: web/controllers/login_controller.go
package controllers

import (
	"fmt"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
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

var loginStaticView = mvc.View{
	Name: "login/login.html",
	Data: iris.Map{"Title": "User Login"},
}

// GetLogin handles GET: http://localhost:8080/login.
func (c *LoginController) Get() mvc.Result {
	return loginStaticView
}

// PostLogin handles POST: http://localhost:8080/login.
func (c *LoginController) Post() mvc.Result {
	var (
		username = c.Ctx.FormValue("UserName")
		password = c.Ctx.FormValue("Password")
	)

	if username == "hptbee" && password == "10031993" {
		var (
			cookieNameForSessionID = "hptcookie"
			sess                   = sessions.New(
				sessions.Config{
					Cookie:  cookieNameForSessionID,
					Expires: time.Hour * 2,
				})
		)
		fmt.Println(c.Session)
		// session2 := c.Session.Increment(sessionIDKey, 23)
		// fmt.Println(session2)
		session := sess.Start(c.Ctx)
		fmt.Println(session)

		session.Set("userid", 23)
		fmt.Println(session)
		return mvc.Response{
			Path: "/user/profile",
		}
	}

	return mvc.Response{
		Path: "/login",
	}
}
