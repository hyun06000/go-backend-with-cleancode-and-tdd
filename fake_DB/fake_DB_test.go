package fakedb

import (
	"testing"
)

func TestShowDataBase(t *testing.T) {

	got := initDBAndSendQuery("show database")
	want := DBMessage{
		Error:     "None",
		Terminal:  "[]",
		CurDB:     "",
		LenDBList: 0,
	}
	checkDBMessage(t, got, want)
}

func TestCreateDataBase(t *testing.T) {

	got := initDBAndSendQuery("CREATE DATABASE FakeDB")
	want := DBMessage{
		Error:     "None",
		Terminal:  "[FakeDB]",
		CurDB:     "",
		LenDBList: 1,
	}
	checkDBMessage(t, got, want)
}

func TestCreateManyDataBases(t *testing.T) {
	fakedbNameList := []string{"A", "B", "C", "D", "E"}
	got := createManyDataBases(fakedbNameList)

	terminal := convertAndSortDBNameList(fakedbNameList)
	want := DBMessage{
		Error:     "None",
		Terminal:  terminal,
		CurDB:     "",
		LenDBList: len(fakedbNameList),
	}
	checkDBMessage(t, got, want)
}

func TestUseDataBase(t *testing.T) {
	fakedbName := "FakeDB"
	got := initDBAndSendQuery(
		"CREATE DATABASE "+fakedbName,
		"use "+fakedbName,
	)

	want := DBMessage{
		Error:     "None",
		Terminal:  "Database changed",
		CurDB:     DBName(fakedbName),
		LenDBList: 1,
	}
	checkDBMessage(t, got, want)
}

func TestCreateTable(t *testing.T) {
	fakedbName := "FakeDB"
	tableName := "TableA"
	tableColumns := "(name string, score string)"

	got := initDBAndSendQuery(
		"CREATE DATABASE "+fakedbName,
		"use "+fakedbName,
		"CREATE TABLE "+tableName+" "+tableColumns,
	)

	want := DBMessage{
		Error:     "None",
		Terminal:  "Table is Created",
		CurDB:     DBName(fakedbName),
		LenDBList: 1,
		CurTable:  TBName(tableName),
		lenTable:  0,
		Columns:   "|name string|score string|",
	}
	checkDBMessage(t, got, want)
}

func TestInsertInto(t *testing.T) {
	fakedbName := "FakeDB"
	tableName := "TableA"
	tableColumns := "(name string, score string)"
	content := "('Jhon', 316)"

	got := initDBAndSendQuery(
		"CREATE DATABASE "+fakedbName,
		"use "+fakedbName,
		"CREATE TABLE "+tableName+" "+tableColumns,
		"INSERT INTO "+tableName+tableColumns+" VALUES "+content,
	)

	columns := convertAndSortTBColumns(tableColumns)
	want := DBMessage{
		Error:     "None",
		Terminal:  content,
		CurDB:     DBName(fakedbName),
		LenDBList: 1,
		CurTable:  TBName(tableName),
		lenTable:  1,
		Columns:   columns,
	}
	checkDBMessage(t, got, want)
}

func initFakeDB() FakeDB {
	fakedb := FakeDB{}
	fakedb.DBMsg.initDBMessage()

	return fakedb
}

func initDBAndSendQuery(querys ...string) DBMessage {
	fakedb := initFakeDB()
	var dbMsg DBMessage
	for _, q := range querys {
		dbMsg = fakedb.Query(q)
	}

	return dbMsg
}

func createManyDataBases(fakedbNameList []string) DBMessage {

	fakedb := initFakeDB()
	var db_msg DBMessage
	for _, fakedbName := range fakedbNameList {
		db_msg = fakedb.Query("CREATE DATABASE " + fakedbName)
	}

	return db_msg
}

func checkDBMessage(t *testing.T, got DBMessage, want DBMessage) {
	assertStringDifference(t, "Error", got.Error, want.Error)
	assertStringDifference(t, "Terminal", got.Terminal, want.Terminal)
	assertStringDifference(t, "DBListString", got.DBListString, want.DBListString)
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

func convertAndSortDBNameList(DBNameList []string) string {
	cvtArgs := GenConvertArgs("SquareBraket_WhiteSpace")
	terminal := ConvertListToString(DBNameList, cvtArgs)
	return terminal
}

func convertAndSortTBColumns(tableColumns string) string {
	cvtArgs := GenConvertArgs("RoundBraket_CommaSpace")
	tableColumnsList := ConvertStringToList(tableColumns, cvtArgs)

	cvtArgs = GenConvertArgs("BarBar_Bar")
	columns := ConvertListToString(tableColumnsList, cvtArgs)

	return columns
}
