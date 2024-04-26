package main

import (
	"log"
	app "middleware/internal/pkg"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Start()
	if err != nil {
		log.Fatal(err)
	}

}
