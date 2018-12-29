package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, "%d", score)

}

func getPlayerScore(player string) int {
	if player == "Pepper" {
		return 20

	}
	if player == "Floyd" {
		return 10

	}
	return 0
}
