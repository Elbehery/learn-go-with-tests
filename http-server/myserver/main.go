package main

import (
	"github.com/quii/learn-go-with-tests/http-server/myserver/app"
	"log"
	"net/http"
)

type InMemStore struct{}

func (s *InMemStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	s := app.PlayerServer{Store: &InMemStore{}}
	handler := http.HandlerFunc(s.ServerHttp)
	log.Fatal(http.ListenAndServe(":5001", handler))
}
