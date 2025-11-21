package main

import (
	"log"

	"github.com/andreyxaxa/unixsort/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Print(err)
	}
}
