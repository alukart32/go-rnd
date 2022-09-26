package main

import "alukart32.com/usage/sql/internal/pg/pglib"

func main() {
	pglib.InsertIntoTestTableAfterCreationInTx()
}
