// file: main.go

package main

import (
	"example.mvc/authen"
	"example.mvc/context"
	"example.mvc/mvc/controllers"
	"example.mvc/repo"
	"example.mvc/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

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
	// Serve our controllers.

	mvc.Configure(app.Party("/hello"), hello)
	mvc.Configure(app.Party("/login"), login)

	mvc.Configure(app.Party("/user"), user)
	// You can also split the code you write to configure an mvc.Application
	// using the `mvc.Configure` method, as shown below.
	mvc.Configure(app.Party("/movies"), movies)

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
	app.Router.Use(authen.BasicAuth)
	app.Handle(new(controllers.HelloController))
}

func user(app *mvc.Application) {
	app.Router.Use(authen.BasicAuth)
	app.Handle(new(controllers.UserController))
}

// note the mvc.Application, it's not iris.Application.
func movies(app *mvc.Application) {
	app.Router.Use(authen.BasicAuth)

	// Create our movie repository with some (memory) data from the datasource.
	repo := repositories.NewMovieRepository(datasource.Movies)
	// Create our movie service, we will bind it to the movie app's dependencies.
	movieService := services.NewMovieService(repo)
	app.Register(movieService)
	app.Handle(new(controllers.MovieController))
}
