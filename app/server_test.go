package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)

		got := resp.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)

		got := resp.Body.String()
		want := "10"
		assertResponseBody(t, got, want)

	})
}

func newGetScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
