package main

import (
	"log"
	"net/http"

	"github.com/gpifko/go-specs-greet/adapters/httpserver"
)

func main() {
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", httpserver.NewHandler()))
}
