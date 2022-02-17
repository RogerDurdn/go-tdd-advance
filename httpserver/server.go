package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case http.MethodGet:
		s.showScore(w,r)
	case http.MethodPost:
		s.processWin(w,r)
	}
}

func (s *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	playerScore := s.store.GetPlayerScore(playerName)
	if playerScore == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	}else{
		fmt.Fprint(w, playerScore)
	}
}
func (s *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	 s.store.RecordWin(playerName)
	w.WriteHeader(http.StatusAccepted)
}

