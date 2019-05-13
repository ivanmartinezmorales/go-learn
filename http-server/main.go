package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore - stores player data in memory
type InMemoryPlayerStore struct{}

// GetPlayerScore - Returns the player score from the InMemoryStore
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not listen on port 500 %v", err)
	}
}
