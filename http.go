package main

import (
	"fmt"
	"net/http"
	"strings"
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

func main() {}
