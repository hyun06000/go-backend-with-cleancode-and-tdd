package myhttp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hyun06000/go-backend-with-cleancode-and-tdd/fakeDB"
)

const dbName string = "PlayerScoreDB"
const tbName string = "GameA"
const cols string = "(player string, score int)"

func createDB() fakeDB.FakeDB {

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

	query := GenQueryToGetScoreWithPlayer(tbName, "'Lee'")
	got := f.Query(query).SelectedValue

	want := 60
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

const prefixURLPlayer string = "/players/"

func TestGetPlayerScoerFromDB(t *testing.T) {
	f := createDB()

	player := "'Lee'"
	score := 60

	want := PlayerScore{
		Player: player,
		Score:  score,
	}
	url := "/" + tbName + prefixURLPlayer + player

	GET(t, f, url, want)
}

func TestGenQueryToUpdateScoreWithPlayer(t *testing.T) {
	f := createDB()

	player := "'Oh'"
	score := 10

	query := GenQueryToUpdateScoreWithPlayer(tbName, cols, player, score)
	f.Query(query)

	query = GenQueryToGetScoreWithPlayer(tbName, player)
	got := f.Query(query).SelectedValue

	if got != score {
		t.Errorf("got %q want %q", got, score)
	}

}

func TestPostPlayerScoerFromDB(t *testing.T) {
	f := createDB()

	player := "'Oh'"
	score := 10

	newPlayerScore := PlayerScore{
		Player: player,
		Score:  score,
	}

	url := "/" + tbName + prefixURLPlayer + player
	POST(t, f, url, newPlayerScore)
	GET(t, f, url, newPlayerScore)
}

func POST(t *testing.T, f fakeDB.FakeDB, url string, want PlayerScore) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&want)

	request, _ := http.NewRequest(http.MethodPost, url, buf)

	responseRecorder := httptest.NewRecorder()

	PlayerScoerUnit(responseRecorder, request, true, f)
}

func GET(
	t *testing.T, f fakeDB.FakeDB, url string, want PlayerScore) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	responseRecorder := httptest.NewRecorder()

	PlayerScoerUnit(responseRecorder, request, true, f)

	got := PlayerScore{}
	json.NewDecoder(responseRecorder.Body).Decode(&got)

	assertPlayerScore(t, got, want)
}

func assertPlayerScore(t *testing.T, got PlayerScore, want PlayerScore) {
	if got.Player != want.Player {
		t.Errorf("got %q want %q", got.Player, want.Player)
	}
	if got.Score != want.Score {
		t.Errorf("got %d want %d", got.Score, want.Score)
	}

}
