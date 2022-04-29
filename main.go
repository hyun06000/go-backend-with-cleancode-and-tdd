package main

import (
	"fmt"

	"github.com/hyun06000/go-backend-with-cleancode-and-tdd/fakeDB"
)

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
