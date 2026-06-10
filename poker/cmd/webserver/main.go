package main

import (
	"log"
	"net/http"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatal("problem starting server: ", err)
	}

	log.Fatalf("could not listen on port 8080 %v", http.ListenAndServe(":8080", server))
}
