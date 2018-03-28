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
	needAuth := app.Party("/user")
	{
		//http://localhost:8080/user/login
		needAuth.Get("/login", func(ctx iris.Context) {})
		// http://localhost:8080/user/profile
		needAuth.Get("/profile", authen.BasicAuth)

	}
	mvc.Configure(needAuth, user)
	// You can also split the code you write to configure an mvc.Application
	// using the `mvc.Configure` method, as shown below.
	mvc.Configure(app.Party("/movies"), movies)

	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/user/login") })

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

func hello(app *mvc.Application) {
	app.Router.Use(authen.BasicAuth)
	app.Handle(new(controllers.HelloController))
}

func user(app *mvc.Application) {

	app.Handle(new(controllers.UserController))
}

// note the mvc.Application, it's not iris.Application.
func movies(app *mvc.Application) {
	// Add the basic authentication(admin:password) middleware
	// for the /movies based requests.
	app.Router.Use(authen.BasicAuth)

	// Create our movie repository with some (memory) data from the datasource.
	repo := repositories.NewMovieRepository(datasource.Movies)
	// Create our movie service, we will bind it to the movie app's dependencies.
	movieService := services.NewMovieService(repo)
	app.Register(movieService)

	// serve our movies controller.
	// Note that you can serve more than one controller
	// you can also create child mvc apps using the `movies.Party(relativePath)` or `movies.Clone(app.Party(...))`
	// if you want.
	app.Handle(new(controllers.MovieController))
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	// third parameter it will be always true because the middleware
	// makes sure for that, otherwise this handler will not be executed.

	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}
