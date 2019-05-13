package main

import (
	"fmt"
	"net/http"
)

// StubPlayerStore - holds the scores
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

// GetPlayerScore - Returns the player's score from the store dictionary
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

// RecordWin - records the wins that the player has
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// PlayerStore - an Interface with all the methods of PlayerScore
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer - a server to house our player
type PlayerServer struct {
	store PlayerStore
}

// ServeHTTP - serves HTTP GET method
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

// GetPlayerScore - gets the score of the player
func GetPlayerScore(name string) string {
	if name == "Pepe" {
		return "20"
	}

	if name == "Phil" {
		return "10"
	}

	return ""
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}
