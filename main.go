// file: main.go

package main

import (
	"fmt"

	"example.mvc/authen"
	"example.mvc/mvc/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

//Check authen of session
func CheckAuthen(ctx *authen.Context) {

	var sesss = ctx.Session()
	var userID = sesss.GetStringDefault("userid", "")
	fmt.Println(userID)
	fmt.Println(sesss)
	// Check if user is authenticated
	if userID != "" {
		ctx.Next()
	}

	ctx.Redirect("/login")
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

	mvc.Configure(app.Party("/login"), login)
	mvc.Configure(app.Party("/user"), user)
	// You can also split the code you write to configure an mvc.Application
	// using the `mvc.Configure` method, as shown below.
	todo := app.Party("/todo")
	mvc.Configure(todo, todos)

	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/todo") })

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

func user(app *mvc.Application) {
	app.Router.Use(authen.Handler(CheckAuthen))
	app.Handle(new(controllers.UserController))
}

func todos(app *mvc.Application) {
	app.Router.Use(authen.Handler(CheckAuthen))
	app.Handle(new(controllers.TodoController))
}
