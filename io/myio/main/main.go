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

	playerServer := myio.NewPlayerServer(&myio.FileSystemPlayerStore{db})
	if err := http.ListenAndServe(":5003", playerServer); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
