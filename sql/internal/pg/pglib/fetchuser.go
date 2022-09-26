package pglib

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"alukart32.com/usage/sql/internal/conf"
	_ "github.com/lib/pq"
)

func FetchUserQuery() {
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

	var (
		id   int
		name string
	)

	rows, err := db.Query("select id, name from public.user where name like 'U%'")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nquery results...")
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}
	if err = rows.Close(); err != nil {
		log.Fatal(err)
	}
}

func PrepareQueryFetch() {
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

	var (
		id   int
		name string
	)

	stmn, err := db.Prepare("select id, name from public.user where name like $1 or name like $2")
	if err != nil {
		log.Fatal(err)
	}
	defer stmn.Close()

	rows, err := stmn.Query("U%", "_v%")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nquery results...")
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}
	if err = rows.Close(); err != nil {
		log.Fatal(err)
	}
}

func SingleQueryFetch() {
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

	var (
		id   int
		name string
	)

	row := db.QueryRow("select id, name from public.user where name like $1", "U%")
	if err = row.Scan(&id, &name); errors.Is(err, sql.ErrNoRows) {
		log.Printf("no row")
	} else {
		log.Fatal(err)
	}

	fmt.Println("\nquery results...")
	log.Println(id, name)
}
