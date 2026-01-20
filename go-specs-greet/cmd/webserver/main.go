package main

import (
	"log"
	"net/http"

	"github.com/gpifko/go-specs-greet/adapters/webserver"
)

func main() {
	handler, err := webserver.NewHandler()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", handler))
}
