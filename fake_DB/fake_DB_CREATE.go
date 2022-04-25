package fakedb

import (
	"strings"
)

func CREATE(f *FakeDB, query string) DBMessage {
	query = strings.TrimPrefix(query, "CREATE ")

	switch {
	case StartWith_DATABASE(query):
		return CREATE_DATABASE(f, query)

	case StartWith_TABLE(query):
		return CREATE_TABLE(f, query)

	default:
		return f.DBMsg
	}
}

func CREATE_DATABASE(f *FakeDB, query string) DBMessage {
	query = strings.TrimPrefix(query, "DATABASE ")

	f.CurrentDB = DBName(query)
	f.initDataBase()

	DBList := f.getDBList()
	cvtArgs := GenConvertArgs("SquareBraket_WhiteSpace")
	f.DBMsg.Terminal = ConvertListToString(DBList, cvtArgs)
	f.DBMsg.LenDBList = len(DBList)

	return f.DBMsg
}

func CREATE_TABLE(f *FakeDB, query string) DBMessage {
	query = strings.TrimPrefix(query, "TABLE ")

	newTableNameString := GetSplitedWord(query, 0)
	newTableName := TBName(newTableNameString)
	f.initTable(newTableName)
	table := f.MySQL[f.CurrentDB][newTableName]

	colsString := TrimPreSuf(query, newTableNameString+" (", ")")
	colsList := Split(colsString, ", ")

	lastColName := f.setColumnsAndTypes(table, colsList)
	colNameList, typeLsit := f.getColumnAndTypeList(table)
	columnsSpec := MergeWordList(colNameList, typeLsit, " ")

	f.DBMsg.Terminal = "Table is Created"
	f.DBMsg.LenDBList = len(f.MySQL)
	f.DBMsg.CurTable = newTableName
	f.DBMsg.lenTable = len(*table[lastColName].contentList)
	cvtArgs := GenConvertArgs("BarBar_Bar")
	f.DBMsg.Columns = ConvertListToString(columnsSpec, cvtArgs)

	return f.DBMsg
}

func (f *FakeDB) initDataBase() {
	if f.MySQL[f.CurrentDB] == nil {
		f.MySQL[f.CurrentDB] = make(DataBase)
	}
}

func (f *FakeDB) initTable(newTableName TBName) {
	if f.MySQL[f.CurrentDB][newTableName] == nil {
		f.MySQL[f.CurrentDB][newTableName] = make(Table)
	}
}

func (f *FakeDB) getColumnAndTypeList(tb Table) ([]string, []string) {
	ColumnNameList := []string{}
	ColumnTypeList := []string{}
	for colName, contents := range tb {
		ColumnNameList = append(ColumnNameList, string(colName))
		ColumnTypeList = append(ColumnTypeList, contents.Type)
	}
	return ColumnNameList, ColumnTypeList
}

func (f *FakeDB) setColumnsAndTypes(table Table, colsList []string) ColumnName {
	var lastColumnName ColumnName
	for _, colAndType := range colsList {
		wordList := Split(colAndType, " ")
		lastColumnName = ColumnName(wordList[0])
		colType := wordList[1]
		table[lastColumnName] = Column{
			Type:        colType,
			contentList: &[]string{},
		}
	}
	return lastColumnName
}
