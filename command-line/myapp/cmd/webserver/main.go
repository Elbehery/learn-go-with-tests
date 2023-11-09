package main

import (
	"github.com/quii/learn-go-with-tests/command-line/myapp"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	f, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file %s: %v", dbFileName, err)
	}

	store, err := myapp.NewFileSystemPlayerStore(f)
	if err != nil {
		log.Fatalf("error creating fileSystemPlayerStore: %v", err)
	}

	server := myapp.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5005", server))
}
