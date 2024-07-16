package main

import (
	"fmt"
	"log"

	"github.com/idkwhyureadthis/test-task/internal/app"
	"github.com/idkwhyureadthis/test-task/internal/pkg/storage"
)

const ADDR = ":8080"

func main() {
	storage.Init()

	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("starting server at http://localhost%s/", ADDR)
	if err := app.Run(ADDR); err != nil {
		log.Fatal(err)
	}
}
