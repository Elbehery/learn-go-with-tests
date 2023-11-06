package main

import (
	"github.com/quii/learn-go-with-tests/io/myio"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening db file %s, %v", dbFileName, err)
	}

	store, err := myio.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store: '%v'", err)
	}
	playerServer := myio.NewPlayerServer(store)
	if err := http.ListenAndServe(":5003", playerServer); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
