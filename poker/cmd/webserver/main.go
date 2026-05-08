package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
)

const dbFileName = "game.db.json"

func main() {
	// server := poker.NewPlayerServer(poker.NewInMemoryPlayerStore()) // in case you need to switch out the file database for the in-memory one
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		log.Fatalf("problem opening %s %v, ", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := poker.NewPlayerServer(store)

	log.Fatalf("could not listen on port 8080 %v", http.ListenAndServe(":8080", server))
}
