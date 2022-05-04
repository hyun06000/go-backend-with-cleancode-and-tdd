package main

import (
	"net/http"

	"github.com/hyun06000/go-backend-with-cleancode-and-tdd/myhttp"
)

func main() {
	http.HandleFunc("/GameA/player", myhttp.PlayerScoer)
	http.ListenAndServe(":8880", nil)
}
