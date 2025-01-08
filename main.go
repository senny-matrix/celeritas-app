package main

import (
	"github.com/senny-matrix/celeritas"
	"github.com/senny-matrix/myapp/data"
	"github.com/senny-matrix/myapp/handlers"
	"github.com/senny-matrix/myapp/middleware"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
