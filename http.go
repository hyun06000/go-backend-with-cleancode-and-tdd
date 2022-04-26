package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hyun06000/go-backend-with-cleancode-and-tdd/fakeDB"
)

var PlayerScoreDB = map[string]string{
	"Pepper": "20",
	"Martin": "16",
	"Trever": "46",
	"Peter":  "52",
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, PlayerScoreDB[playerName])
}

func main() {
	dbMsg := fakeDB.InitDBAndSendQuery(
		"CREATE DATABASE fakedb",
		"USE fakedb",
		"CREATE TABLE tableA (name string, score string)",
		"INSERT INTO tableA(name string, score string) VALUES ('A', 316)",
		"INSERT INTO tableA(name string, score string) VALUES ('B', 521)",
		"SELECT * FROM tableA",
	)

	fmt.Println(dbMsg.Terminal)

}
