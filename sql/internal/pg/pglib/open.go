package pglib

import (
	"database/sql"
	"fmt"

	"alukart32.com/usage/sql/internal/conf"
	_ "github.com/lib/pq"
)

func PgLibSqlOpen() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBname)

	// sql.Open Does not establish any connections to the database
	// prepares the database abstraction for later use.
	// The first actual connection to the underlying datastore
	// will be established lazily, when it’s needed for the first time
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to postgres...")
}
