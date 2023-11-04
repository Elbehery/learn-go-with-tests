package main

import (
	"github.com/quii/learn-go-with-tests/io/myio"
	"log"
	"net/http"
)

func main() {
	playerServer := myio.NewPlayerServer(myio.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5003", playerServer))
}
