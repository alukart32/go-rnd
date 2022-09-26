package pgx

import (
	"context"
	"fmt"
	"log"
	"os"

	"alukart32.com/usage/sql/internal/conf"
	"github.com/jackc/pgx/v4"
)

func PgxSqlOpen() {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DBname)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully connected to postgres...")
}
