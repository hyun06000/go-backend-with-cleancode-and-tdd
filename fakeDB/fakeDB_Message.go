package fakeDB

type DBMessage struct {
	Error         string
	Terminal      string
	CurDB         DBName
	DBListString  string
	LenDBList     int
	CurTable      TBName
	lenTable      int
	Columns       string
	SelectedValue int
}

func (dbmsg *DBMessage) initDBMessage() {

	dbmsg.Error = "None"
	dbmsg.Terminal = ""
	dbmsg.CurDB = ""
	dbmsg.LenDBList = 0
	dbmsg.CurTable = ""
	dbmsg.lenTable = 0
	dbmsg.Columns = ""
	dbmsg.SelectedValue = -1
}
