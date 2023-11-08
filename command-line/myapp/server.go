package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store  PlayerStore
	router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	srvr := &PlayerServer{}
	srvr.store = store

	router := &http.ServeMux{}
	router.Handle("/players/", http.HandlerFunc(srvr.playersHandler))
	router.Handle("/league", http.HandlerFunc(srvr.leagueHandler))

	srvr.router = router
	return srvr
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(s.store.GetLeague())
}

func (s *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		s.showScore(w, player)
	case http.MethodPost:
		s.processWin(w, player)
	}
}

func (s *PlayerServer) showScore(w http.ResponseWriter, name string) {
	score := s.store.GetPlayerScore(name)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) processWin(w http.ResponseWriter, name string) {
	s.store.RecordWin(name)
	w.WriteHeader(http.StatusAccepted)
}
