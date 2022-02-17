package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	recordHits int
}

func (s StubPlayerStore) GetPlayerScore(name string) int  {
	score, ok := s.scores[name]
	if ok {
		return score
	}
	return 0
}

func (s *StubPlayerStore) RecordWin(name string) {
	score := s.GetPlayerScore(name)
	s.scores[name] = score + 1
	s.recordHits++
}

func TestGETPlayers(t *testing.T) {
	stubPlayerStore := &StubPlayerStore{scores: make(map[string]int)}
	stubPlayerStore.scores["roger"] = 20
	stubPlayerStore.scores["juan"] = 10

	playerServer := &PlayerServer{store: stubPlayerStore}

	for k, v := range stubPlayerStore.scores {
	t.Run("returns "+k+" score", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/players/"+k, nil)
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Body.String()
		want := v

		assertStatusCode(t,response.Code, http.StatusOK)
		if got != strconv.Itoa(want) {
			t.Errorf("expected %v, got %v", want, got)
		}
	})
	}
	t.Run("Get 404 by not player found", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/players/noValid", nil)
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Code

		assertStatusCode(t,got, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	stubPlayerStore := StubPlayerStore{scores: make(map[string]int)}
	stubPlayerStore.scores["roger"] = 20
	stubPlayerStore.scores["pepe"] = 20

	t.Run("add win score to previous player", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Roger", nil)
		response := httptest.NewRecorder()
		playerServer := &PlayerServer{&stubPlayerStore}
		playerServer.ServeHTTP(response, request)
		playerServer.ServeHTTP(response, request)

		assertStatusCode(t,response.Code, http.StatusAccepted)
		if stubPlayerStore.recordHits != 2 {
			t.Errorf("expected 1 hits on recordWins but got %d", stubPlayerStore.recordHits)
		}
	})

}

func assertStatusCode(t *testing.T, got, want int)  {
	t.Helper()
	if got != want{
		t.Fatalf("expected %v, but got %v", want, got)
	}
}