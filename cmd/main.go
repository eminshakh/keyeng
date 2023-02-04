package main

import (
	"github.com/eminshakh/keyeng/internal/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
