package fakedb

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

	switch {
	case StartWith_SHOW(query):
		return SHOW(f, query)

	case StartWith_CREATE(query):
		return CREATE(f, query)

	case StartWith_USE(query):
		return USE(f, query)

	case StartWith_INSERT_INTO(query):
		return INSERT_INTO(f, query)

	case StartWith_SELECT(query):
		return SELECT(f, query)

	default:
		return f.DBMsg
	}
}

func (f *FakeDB) initMySQL() {
	if f.MySQL == nil {
		f.MySQL = make(MySQL)
	}
}
