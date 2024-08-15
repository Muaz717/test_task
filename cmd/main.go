package main

import (
	"log"

	"github.com/Muaz717/test_task/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal("failed to init app")
	}

	err = a.Run()
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
