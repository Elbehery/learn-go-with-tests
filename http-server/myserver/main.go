package main

import (
	"github.com/quii/learn-go-with-tests/http-server/myserver/app"
	"log"
	"net/http"
)

func main() {
	s := &app.PlayerServer{}
	handler := http.HandlerFunc(s.ServerHttp)
	log.Fatal(http.ListenAndServe(":5001", handler))
}
