package main

import (
	"log"

	"github.com/nbonair/GoIDM/internal/handlers/http"
)

func main() {
	server := http.NewServer()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
