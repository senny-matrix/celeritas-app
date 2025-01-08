package main

import (
	"github.com/senny-matrix/celeritas"
	"github.com/senny-matrix/myapp/data"
	"github.com/senny-matrix/myapp/handlers"
	"github.com/senny-matrix/myapp/middleware"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatalln(err)
	}

	cel.AppName = "MyApp"
	myMiddleware := &middleware.Middleware{
		App: cel,
	}
	myHandlers := &handlers.Handlers{
		App: cel,
	}

	app := &application{
		App:        cel,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
