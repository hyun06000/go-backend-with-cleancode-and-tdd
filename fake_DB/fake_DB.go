package fakedb

import (
	"strings"
)

type DBName string
type TBName string
type ColumnName string
type Column struct {
	Type        string
	contentList *[]string
}

type Table map[ColumnName]Column
type DataBase map[TBName]Table
type MySQL map[DBName]DataBase

// FakeDB.MySQL[DBName][TBName][ColumnName].contentList[index]

type FakeDB struct {
	CurrentDB DBName
	MySQL     MySQL
	DBMsg     DBMessage
}

func (f *FakeDB) Query(query string) DBMessage {
	f.initMySQL()

	if IsShowDataBase(query) {
		cvtArgs := GenConvertArgs("SquareBraket_WhiteSpace")
		f.DBMsg.Terminal = ConvertListToString(f.getDBList(), cvtArgs)
		return f.DBMsg

	} else {
		if StartWith_CREATE(query) {
			query = strings.TrimPrefix(query, "CREATE ")

			if StartWith_DATABASE(query) {
				query = strings.TrimPrefix(query, "DATABASE ")

				f.CurrentDB = DBName(query)
				f.initDataBase()

				DBList := f.getDBList()
				cvtArgs := GenConvertArgs("SquareBraket_WhiteSpace")
				f.DBMsg.Terminal = ConvertListToString(DBList, cvtArgs)
				f.DBMsg.LenDBList = len(DBList)

			} else if StartWith_TABLE(query) {
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
			}

		} else if StartWith_use(query) {
			query = strings.TrimPrefix(query, "use ")
			f.CurrentDB = DBName(query)

			f.DBMsg.Terminal = "Database changed"
			f.DBMsg.CurDB = f.CurrentDB
			f.DBMsg.LenDBList = len(f.MySQL)

		} else if StartWith_INSERT_INTO(query) {
			query := strings.TrimPrefix(query, "INSERT INTO ")
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

		}
	}
	return f.DBMsg
}

func (f *FakeDB) initMySQL() {
	if f.MySQL == nil {
		f.MySQL = make(MySQL)
	}
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

func (f *FakeDB) getDBList() []string {
	DBList := []string{}
	for key, _ := range f.MySQL {
		DBList = append(DBList, string(key))
	}
	return DBList
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
