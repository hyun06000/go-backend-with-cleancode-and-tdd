package fakeDB

import (
	"strings"
)

func SHOW(f *FakeDB, query string) DBMessage {
	query = strings.TrimPrefix(query, "SHOW ")

	switch {
	case StartWith_DATABASE(query):
		return SHOW_DATABASE(f, query)

	case StartWith_TABLES(query):
		return SHOW_TABLES(f, query)

	default:
		return f.DBMsg
	}
}

func SHOW_DATABASE(f *FakeDB, query string) DBMessage {
	cvtArgs := GenConvertArgs("SquareBraket_WhiteSpace")
	f.DBMsg.Terminal = ConvertListToString(f.getDBList(), cvtArgs)
	return f.DBMsg
}

func SHOW_TABLES(f *FakeDB, query string) DBMessage {
	cvtArgs := GenConvertArgs("SquareBraket_WhiteSpace")
	f.DBMsg.Terminal = ConvertListToString(f.getTBList(), cvtArgs)
	return f.DBMsg
}

func (f *FakeDB) getDBList() []string {
	DBList := []string{}
	for key, _ := range f.MySQL {
		DBList = append(DBList, string(key))
	}
	return DBList
}

func (f *FakeDB) getTBList() []string {
	TBList := []string{}
	for key, _ := range f.MySQL[f.CurrentDB] {
		TBList = append(TBList, string(key))
	}
	return TBList
}
