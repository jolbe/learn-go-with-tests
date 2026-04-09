package main

import (
	"log"
	"net/http"

	httpserver "github.com/gregor-pifko/learn-go-with-tests/http-server"
)

func main() {
	server := httpserver.NewPlayerServer(httpserver.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":8080", server))
}
