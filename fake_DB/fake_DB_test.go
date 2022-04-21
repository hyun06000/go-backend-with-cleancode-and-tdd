package fakedb

import (
	"testing"
)

func TestShowDataBase(t *testing.T) {
	fakedb := initFakeDB()
	got := fakedb.Query("show database")

	want := DBMessage{
		Error:     "None",
		Terminal:  "[]",
		CurDB:     "",
		LenDBList: 0,
	}

	checkDBMessage(t, got, want)
}

func TestCreateDataBase(t *testing.T) {
	fakedb := initFakeDB()
	got := fakedb.Query("CREATE DATABASE FakeDB")

	want := DBMessage{
		Error:     "None",
		Terminal:  "[FakeDB]",
		CurDB:     "",
		LenDBList: 1,
	}
	checkDBMessage(t, got, want)
}

func TestCreateManyDataBases(t *testing.T) {
	fakedb := initFakeDB()
	fakedbNameList := []string{"A", "B", "C", "D", "E"}
	fakedbNameListString, got := createManyDataBases(&fakedb, fakedbNameList)

	want := DBMessage{
		Error:     "None",
		Terminal:  fakedbNameListString,
		CurDB:     "",
		LenDBList: len(fakedbNameList),
	}
	checkDBMessage(t, got, want)
}

func TestUseDataBase(t *testing.T) {
	fakedb := initFakeDB()
	fakedbName := "FakeDB"
	fakedb.Query("CREATE DATABASE " + fakedbName)
	got := fakedb.Query("use " + fakedbName)

	want := DBMessage{
		Error:     "None",
		Terminal:  "Database changed",
		CurDB:     fakedbName,
		LenDBList: 1,
	}
	checkDBMessage(t, got, want)
}

func TestCreateTable(t *testing.T) {
	fakedb := initFakeDB()
	fakedbName := "FakeDB"
	fakedb.Query("CREATE DATABASE " + fakedbName)
	fakedb.Query("use " + fakedbName)

	tableName := "TableA"
	tableColumns := "(name string, score int)"
	got := fakedb.Query("CREATE TABLE " + tableName + " " + tableColumns)

	want := DBMessage{
		Error:     "None",
		Terminal:  "Table is Created",
		CurDB:     fakedbName,
		LenDBList: 1,
		CurTable:  tableName,
		lenTable:  0,
		Columns:   "|name string|score int|",
	}
	checkDBMessage(t, got, want)
}

func initFakeDB() MySQL {
	fakedb := MySQL{}
	fakedb.DBMsg.initDBMessage()

	return fakedb
}

func createManyDataBases(fakedb *MySQL, fakedbNameList []string) (string, DBMessage) {

	var db_msg DBMessage
	for _, fakedbName := range fakedbNameList {
		db_msg = fakedb.Query("CREATE DATABASE " + fakedbName)
	}

	fakedbNameListString := ConvertListToString(fakedbNameList)
	return fakedbNameListString, db_msg
}

func checkDBMessage(t *testing.T, got DBMessage, want DBMessage) {
	assertStringDifference(t, "Error", got.Error, want.Error)
	assertStringDifference(t, "Terminal", got.Terminal, want.Terminal)
	assertStringDifference(t, "CurDB", got.CurDB, want.CurDB)
	assertStringDifference(t, "DBListString", got.DBListString, want.DBListString)
	assertStringDifference(t, "CurTable", got.CurTable, want.CurTable)
	assertStringDifference(t, "Columns", got.Columns, want.Columns)

	assertintDifference(t, "LenDBList", got.LenDBList, want.LenDBList)
	assertintDifference(t, "lenTable", got.lenTable, want.lenTable)
}

func assertStringDifference(t *testing.T, testName string, got string, want string) {
	if got != want {
		t.Errorf(testName+" : got %q want %q", got, want)
	}
}

func assertintDifference(t *testing.T, testName string, got int, want int) {
	if got != want {
		t.Errorf(testName+" : got %q want %q", got, want)
	}
}
