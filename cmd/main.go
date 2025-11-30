package main

import (
	"log"

	"github.com/aykuanysh/creditcard/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
