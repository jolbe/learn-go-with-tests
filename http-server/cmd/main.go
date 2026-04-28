package main

import (
	"log"
	"net/http"
	"os"

	httpserver "github.com/gregor-pifko/learn-go-with-tests/http-server"
)

const dbFileName = "game.db.json"

func main() {
	// server := httpserver.NewPlayerServer(httpserver.NewInMemoryPlayerStore()) // in case you need to switch out the file database for the in-memory one
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		log.Fatalf("problem opening %s %v, ", dbFileName, err)
	}

	store, err := httpserver.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := httpserver.NewPlayerServer(store)

	log.Fatalf("could not listen on port 8080 %v", http.ListenAndServe(":8080", server))
}
