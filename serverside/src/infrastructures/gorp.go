package infrastructures

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
	"gopkg.in/gorp.v1"
)

func InitDB() *gorp.DbMap {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Panicf("SQL接続失敗 %s", errors.WithStack(err))
	}

	// // create the table. in a production system you'd generally
	// // use a migration tool, or create the tables via scripts
	// err = dbmap.CreateTablesIfNotExists()
	// checkErr(err, "Create tables failed")

	return &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
}
