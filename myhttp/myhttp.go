package myhttp

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyun06000/go-backend-with-cleancode-and-tdd/fakeDB"
)

type PlayerScore struct {
	Player string
	Score  int
}

func GenQueryToGetScoreWithPlayer(tbName string, player string) string {
	return "SELECT score FROM " + tbName + " WHERE player = " + player
}

func GenQueryToUpdateScoreWithPlayer(
	tbName string, cols string, player string, score int) string {

	insertPrefix := "INSERT INTO " + tbName + cols + " VALUES "
	scoreString := strconv.Itoa(score)

	return insertPrefix + "(" + player + ", " + scoreString + ")"
}

func PlayerScoer(w http.ResponseWriter, r *http.Request) {
	PlayerScoerUnit(w, r, false, fakeDB.FakeDB{})
}

func PlayerScoerUnit(w http.ResponseWriter, r *http.Request, unitTest bool, f fakeDB.FakeDB) {

	switch {
	case r.Method == http.MethodGet:
		url := strings.TrimPrefix(r.URL.Path, "/")
		urlParsingList := fakeDB.Split(url, "/")
		tbName := urlParsingList[0]
		player := urlParsingList[2]

		rtn := PlayerScore{}
		rtn.Player = player

		query := GenQueryToGetScoreWithPlayer(tbName, player)

		if unitTest {
			rtn.Score = f.Query(query).SelectedValue
			json.NewEncoder(w).Encode(rtn)
		}

	case r.Method == http.MethodPost:

		url := strings.TrimPrefix(r.URL.Path, "/")
		urlParsingList := fakeDB.Split(url, "/")
		tbName := urlParsingList[0]

		cols := "(player string, score int)"
		body := PlayerScore{}
		json.NewDecoder(r.Body).Decode(&body)
		query := GenQueryToUpdateScoreWithPlayer(
			tbName, cols, body.Player, body.Score,
		)
		if unitTest {
			f.Query(query)
		}

	}

}
