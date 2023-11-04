package app

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (s *PlayerServer) ServerHttp(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/player/")
	fmt.Fprint(w, s.Store.GetPlayerScore(player))
}
