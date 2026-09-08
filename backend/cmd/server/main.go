package main

import (
	"escoria/internal/bootstrap"
	"log"
)

func main() {
	app, err := bootstrap.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
