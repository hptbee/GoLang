// file: main.go

package main

import (
	"time"

	"example.mvc/mvc/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

const cookieNameForSessionID = "hptcookie"

var sess = sessions.New(sessions.Config{
	Cookie:  cookieNameForSessionID,
	Expires: time.Hour * 2,
})

//Check authen of session
func CheckAuthen(ctx iris.Context) {

	var sesss = sess.Start(ctx)
	var userID = sesss.GetIntDefault("userid", 0)

	// Check if user is authenticated
	if userID == 0 {
		ctx.Redirect("/login")
	}

	// Print secret message
}

// Start Session
func StartSession(ctx iris.Context) {
	sess.Start(ctx)
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Load the template files.
	tmpl := iris.HTML("./mvc/views", ".html").
		Layout("shared/_layout.html").
		Reload(true)
	app.RegisterView(tmpl)

	app.StaticWeb("/content", "./mvc/content")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})

	mvc.Configure(app.Party("/hello"), hello)
	mvc.Configure(app.Party("/login"), login).Register(sess.Start)

	mvc.Configure(app.Party("/user"), user)
	// You can also split the code you write to configure an mvc.Application
	// using the `mvc.Configure` method, as shown below.
	mvc.Configure(app.Party("/todos"), todos)

	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/login") })

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

func login(app *mvc.Application) {

	app.Handle(new(controllers.LoginController))
}
func hello(app *mvc.Application) {
	app.Router.Use(CheckAuthen)
	app.Handle(new(controllers.HelloController))
}

func user(app *mvc.Application) {
	app.Router.Use(CheckAuthen)
	app.Handle(new(controllers.UserController))
}

// note the mvc.Application, it's not iris.Application.
func todos(app *mvc.Application) {
	//app.Router.Use(CheckAuthen)

	// Create our todo service, we will bind it to the todo app's dependencies.
	app.Handle(new(controllers.TodoController))
}
