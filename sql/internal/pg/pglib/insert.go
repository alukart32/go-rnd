package pglib

import (
	"database/sql"
	"fmt"
	"log"

	"alukart32.com/usage/sql/internal/conf"
)

func ExecInsert() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connected to postgres...")

	id, name := 5, "Yan"
	_, err = db.Exec("insert into public.user(id, name) values($1, $2)", id, name)
	if err != nil {
		log.Fatal(err)
	}
}

var createTable = `
CREATE TABLE IF NOT EXISTS public."test"
(
    id integer NOT NULL,
    name character varying(20) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT test_pkey PRIMARY KEY (id)
)
`

func InsertIntoTestTableAfterCreationInTx() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBname)

	var err error
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connected to postgres...")

	// start a new tx to create table "test"
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("start a new tx-1 to postgres...")
	_, err = tx.Exec(createTable)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
	fmt.Println("close tx-1 to postgres...")

	_, err = db.Exec("insert into public.test(id, name) values($1, $2)", 1, "ttt")
	if err != nil {
		log.Fatal(err)
	}
}
