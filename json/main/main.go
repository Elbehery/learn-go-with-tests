package main

import (
	"github.com/quii/learn-go-with-tests/json/myjson"
	"log"
	"net/http"
)

func main() {
	store := myjson.NewInMemoryPlayerStore()
	srvr := myjson.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe("5002", srvr))
}
