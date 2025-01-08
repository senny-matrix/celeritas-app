package data

import (
	"database/sql"
	"fmt"
	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
	"os"
)

var db *sql.DB
var upper db2.Session

type Models struct {
	// Any models inserted here (and New function
	// are easily accessible throughout entire application
}

func New(databasePool *sql.DB) Models {
	db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		upper, _ = mysql.New(databasePool)

	case "postgres", "postgresql":
		upper, _ = postgresql.New(databasePool)

	default:
		// Does nothing
	}

	return Models{}
}

func getInsertID(i db2.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	} else if idType == "uint64" {
		return int(i.(uint64))
	} else {
		return i.(int)
	}
}
