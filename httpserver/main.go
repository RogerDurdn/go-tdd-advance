package main

import (
	"log"
	"net/http"
)

type DefaultPlayerStore struct {
	scores map[string]int
}

func (d DefaultPlayerStore) GetPlayerScore(name string) int {
	score, ok := d.scores[name]
	if ok {
		return score
	}
	return 0
}

func (d DefaultPlayerStore) RecordWin(name string)  {

}

func BasicPlayerStore() *DefaultPlayerStore {
	defaultStore := DefaultPlayerStore{scores: make(map[string]int)}
	defaultStore.scores["roger"] = 20
	defaultStore.scores["juan"] = 10
	return &defaultStore
}

func main()  {
	playerServer := &PlayerServer{store: BasicPlayerStore()}
	handler := http.HandlerFunc(playerServer.ServeHTTP)
	log.Println(http.ListenAndServe(":4000", handler))
}

