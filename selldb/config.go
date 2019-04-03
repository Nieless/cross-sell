package selldb

import (
	"github.com/jmoiron/sqlx"
	"os"
)

var db *sqlx.DB
var DBName = os.Getenv("DB_NAME")   //crosssell

// SetDB sets the global db variable makes
// so that a pool of database connections is
// globally available to the package.  It is
// best practice to initiate db connections from
// a main package and distribute them to libararies
// as needed, so that is what this function does.
func SetDB(dbParam *sqlx.DB) error {
	if err := dbParam.Ping(); err != nil {
		return err
	}

	db = dbParam
	return nil
}