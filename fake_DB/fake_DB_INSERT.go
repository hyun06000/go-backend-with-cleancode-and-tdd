package fakedb

import (
	"strings"
)

func INSERT_INTO(f *FakeDB, query string) DBMessage {
	query = strings.TrimPrefix(query, "INSERT INTO ")

	qListWithoutPrefix := strings.Split(query, " VALUES ")

	tBNameAndCols := qListWithoutPrefix[0]
	tableName, colNameList := GetTBNameAndColNameList(tBNameAndCols)

	contentString := strings.Trim(qListWithoutPrefix[1], "()")
	contentList := Split(contentString, ", ")

	table := f.MySQL[f.CurrentDB][tableName]
	var colName string
	for i, colNameAndType := range colNameList {
		colName = Split(string(colNameAndType), " ")[0]

		tbCol := table[ColumnName(colName)]
		*tbCol.contentList = append(*tbCol.contentList, contentList[i])
	}

	f.DBMsg.Terminal = "(" + contentString + ")"
	f.DBMsg.CurTable = tableName
	f.DBMsg.lenTable = len(*table[ColumnName(colName)].contentList)

	return f.DBMsg
}

func GetTBNameAndColNameList(query string) (TBName, []ColumnName) {
	tbNameAndColNameList := strings.Split(query, "(")

	tbName := TBName(tbNameAndColNameList[0])

	colNames := strings.TrimSuffix(tbNameAndColNameList[1], ")")
	colNameList := Split(colNames, ", ")

	var rtnColNameList []ColumnName
	for _, colName := range colNameList {
		rtnColNameList = append(rtnColNameList, ColumnName(colName))
	}

	return tbName, rtnColNameList

}
