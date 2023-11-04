package main

import (
	"github.com/quii/learn-go-with-tests/http-server/myserver/app"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(app.PlayerServer)
	log.Fatal(http.ListenAndServe(":5001", handler))
}
