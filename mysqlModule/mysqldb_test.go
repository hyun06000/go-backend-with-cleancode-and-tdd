package mysqlmodule

import (
	"fmt"
	"testing"
)

const dbName string = "PlayerScoreDB"
const tbName string = "PlayerScore"

func TestOpenDBObject(t *testing.T) {
	db, err := OpenDBObject(dbName)
	assertOpenError(t, err)
	defer db.Close()
}

func TestCreateTable(t *testing.T) {
	db, err := OpenDBObject(dbName)
	assertOpenError(t, err)
	assertOpenError(t, err)
	err = CreateTable(db, tbName)
	if err != nil {
		if err.Error() == "already exists" {
			_, err = db.Exec("DROP TABLE " + tbName)
			LogError(err)
			err = CreateTable(db, tbName)
		}
	}

	assertError(t, err, "cannot drop table")
	defer db.Close()

}

func TestDropTable(t *testing.T) {
	db, err := OpenDBObject(dbName)
	assertOpenError(t, err)
	err = DropTable(db, tbName)

	if err != nil {
		if err.Error() == "unknown table" {
			q := "CREATE TABLE %s (player TINYTEXT, score INT)"
			_, err = db.Exec(fmt.Sprintf(q, tbName))
			LogError(err)
			err = DropTable(db, tbName)
		}
	}

	assertError(t, err, "cannot drop table")
	defer db.Close()

}

func TestIsExistTable(t *testing.T) {
	db, err := OpenDBObject(dbName)
	assertOpenError(t, err)
	if IsExistTable(db, dbName, tbName) {
		err := DropTable(db, tbName)
		assertOpenError(t, err)
		if IsExistTable(db, dbName, tbName) {
			t.Errorf("func IsExistTable error")
		}
	} else {
		err := CreateTable(db, tbName)
		assertOpenError(t, err)
		if !IsExistTable(db, dbName, tbName) {
			t.Errorf("func IsExistTable error")
		}
	}
	defer db.Close()

}

func TestInitTable(t *testing.T) {
	db, err := OpenDBObject(dbName)
	assertOpenError(t, err)

	err = InitTable(db, dbName, tbName)
	if err != nil {
		if err.Error() == "table is not exist" {
			err = CreateTable(db, tbName)
			LogError(err)
			err = InitTable(db, dbName, tbName)
		}
	}
	assertError(t, err, "cannot init table")

	defer db.Close()

}

func TestInsert(t *testing.T) {
	db, err := OpenDBObject(dbName)
	assertOpenError(t, err)

	err = Insert(db, qInsert, "Choi", 200)
	assertError(t, err, "Cannot send query")
	defer db.Close()

}

func assertOpenError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Cannot open DB")
	}
}

func assertError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf(msg)
	}
}
