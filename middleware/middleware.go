package middleware

import (
	"github.com/senny-matrix/celeritas"
	"github.com/senny-matrix/myapp/data"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
