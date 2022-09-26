package main

import (
	"log"

	"alukart32.com/usage/context/config"
	"alukart32.com/usage/context/internal/app"
)

func main() {
	if err := app.Run(*config.New()); err != nil {
		log.Fatal(err)
	}
}
