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

	fmt.Fprintf(w, "%d", p.store.GetPlayerScore(player))

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
