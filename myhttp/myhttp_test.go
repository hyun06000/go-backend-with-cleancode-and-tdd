package myhttp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hyun06000/go-backend-with-cleancode-and-tdd/fakeDB"
)

func createDB() fakeDB.FakeDB {
	dbName := "PlayerScoreDB"
	tbName := "GameA"
	cols := "(player string, score int)"
	f := fakeDB.InitFakeDB()
	f.Query("CREATE DATABASE " + dbName)
	f.Query("USE " + dbName)
	f.Query("CREATE TABLE " + tbName + " " + cols)
	insertPrefix := "INSERT INTO " + tbName + cols + " VALUES "
	f.Query(insertPrefix + "('Park', 100)")
	f.Query(insertPrefix + "('Kim', 80)")
	f.Query(insertPrefix + "('Lee', 60)")
	f.Query(insertPrefix + "('Choi', 40)")

	return f

}

func TestGenQueryToGetScoreWithPlayer(t *testing.T) {
	f := createDB()

	query := GenQueryToGetScoreWithPlayer("GameA", "'Lee'")
	got := f.Query(query).SelectedValue

	want := 60
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

const prefixURLPlayer string = "/players/"

func TestGetPlayerScoerFromDB(t *testing.T) {
	f := createDB()

	playerName := "'Lee'"
	request, _ := http.NewRequest(http.MethodGet, prefixURLPlayer+playerName, nil)
	responseRecorder := httptest.NewRecorder()

	PlayerScoer(responseRecorder, request, true, f)

	got := PlayerScore{}
	json.NewDecoder(responseRecorder.Body).Decode(&got)
	want := 60
	if got.Player != playerName {
		t.Errorf("got %q want %q", got.Player, playerName)
	}
	if got.Score != want {
		t.Errorf("got %d want %d", got.Score, want)
	}
}
