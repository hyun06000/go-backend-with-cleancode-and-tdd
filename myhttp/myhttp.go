package myhttp

import (
	"encoding/json"
	"net/http"
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

func PlayerScoer(w http.ResponseWriter, r *http.Request, unitTest bool, f fakeDB.FakeDB) {
	rtn := PlayerScore{}
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	rtn.Player = player

	query := GenQueryToGetScoreWithPlayer("GameA", player)

	if unitTest {
		rtn.Score = f.Query(query).SelectedValue
		json.NewEncoder(w).Encode(rtn)
	}
}
