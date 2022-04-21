package fakedb

import "strings"

type MySQL struct {
	DBList    []string
	CurrentDB string
	TBList    map[string][]string
	CurrentTB string
	DBMsg     DBMessage
}

type Table struct {
	Name    string
	Content map[string]map[string]string
}

type DBMessage struct {
	Error        string
	Terminal     string
	CurDB        string
	DBListString string
	LenDBList    int
	CurTable     string
	lenTable     int
	Columns      string
}

func (m *MySQL) Query(query string) DBMessage {
	if query == "show database" {
		m.DBMsg.Terminal = ConvertListToString(m.DBList)
		return m.DBMsg
	} else {
		var parsed []string = strings.Split(query, " ")
		if parsed[0] == "CREATE" {
			if parsed[1] == "DATABASE" {
				m.DBList = append(m.DBList, parsed[2])

				m.DBMsg.Terminal = ConvertListToString(m.DBList)
				m.DBMsg.LenDBList = len(m.DBList)
			} else if parsed[1] == "TABLE" {
				m.CurrentTB = parsed[2]
				curTBList := m.TBList[m.CurrentDB]
				curTBList = append(curTBList, m.CurrentTB)

				m.DBMsg.Terminal = "Table is Created"
				m.DBMsg.LenDBList = len(m.DBList)
				m.DBMsg.CurTable = m.CurrentTB
				m.DBMsg.lenTable = len(curTBList)

				start := strings.Index(query, "(")
				end := strings.Index(query, ")")
				cols := strings.Split(query[start+1:end], ", ")
				cols_rtn := "|"
				for _, col := range cols {
					parsedCol := strings.Split(col, " ")
					colName := parsedCol[0]
					colType := parsedCol[1]
					cols_rtn += colName + " " + colType + "|"
				}
				m.DBMsg.Columns = cols_rtn
				m.DBMsg.lenTable = 0 //ToDo : max(Table index) + 1
			}

		} else if parsed[0] == "use" {
			m.CurrentDB = parsed[1]

			m.DBMsg.Terminal = "Database changed"
			m.DBMsg.CurDB = m.CurrentDB
			m.DBMsg.LenDBList = len(m.DBList)
		}
	}
	return m.DBMsg
}

func (dbmsg *DBMessage) initDBMessage() {

	dbmsg.Error = "None"
	dbmsg.Terminal = ""
	dbmsg.CurDB = ""
	dbmsg.LenDBList = 0
	dbmsg.CurTable = ""
	dbmsg.lenTable = 0
	dbmsg.Columns = ""
}

func ConvertListToString(NameList []string) string {
	ListString := "["
	for i, Name := range NameList {
		ListString += Name
		if i != len(NameList)-1 {
			ListString += " "
		}
	}
	ListString += "]"

	return ListString
}
