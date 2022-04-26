package fakeDB

import (
	"strings"
)

func USE(f *FakeDB, query string) DBMessage {
	query = strings.TrimPrefix(query, "USE ")
	f.CurrentDB = DBName(query)

	f.DBMsg.Terminal = "Database changed"
	f.DBMsg.CurDB = f.CurrentDB
	f.DBMsg.LenDBList = len(f.MySQL)

	return f.DBMsg
}
