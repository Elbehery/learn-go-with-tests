package main

import (
	"github.com/quii/learn-go-with-tests/command-line/myapp"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	f, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening database file %s: %v", dbFileName, err)
	}

	store, err := myapp.NewFileSystemPlayerStore(f)
	if err != nil {
		log.Fatalf("error creating FileSystemPlayerStore: %v", err)
	}

	cli := myapp.NewCLI(store, os.Stdin)
	cli.PlayPoker()
}
