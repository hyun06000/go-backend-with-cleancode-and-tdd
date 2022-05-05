package mysqlmodule

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const qInsert string = "INSERT INTO PlayerScore VALUES (?, ?)"
const qSelect string = "SELECT %s FROM PlayerScore WHERE %s = ?"

func OpenDBObject(dbName string) (*sql.DB, error) {
	driverName := "mysql"
	sqlUserName := "root"
	sqlUserPassWD := 1234
	tcpIp := "127.0.0.1"
	port := 3306

	connectString := fmt.Sprintf("%s:%d@tcp(%s:%d)/%s",
		sqlUserName,
		sqlUserPassWD,
		tcpIp,
		port,
		dbName,
	)

	db, err := sql.Open(driverName, connectString)
	LogError(err)

	err = db.Ping()

	return db, err

}

func IsExistTable(db *sql.DB, dbName string, tbName string) bool {
	var count string
	q := "SELECT COUNT(*) FROM information_schema.tables "
	q += "WHERE table_schema = ? "
	q += "AND table_name = ?"

	db.QueryRow(q, dbName, tbName).Scan(&count)

	return count == "1"
}

func CreateTable(db *sql.DB, tbName string) error {
	var err error
	q := "CREATE TABLE %s (player TINYTEXT, score INT)"
	result, err := db.Exec(fmt.Sprintf(q, tbName))

	if err != nil {
		errMsg := strings.Split(err.Error(), ": ")
		if errMsg[len(errMsg)-1] == fmt.Sprintf("Table '%s' already exists", tbName) {
			return fmt.Errorf("already exists")
		}
		return err
	}

	n, err := result.RowsAffected()
	if n != 0 {
		return fmt.Errorf("CREATE ERROR")
	}

	return err
}

func DropTable(db *sql.DB, tbName string) error {
	result, err := db.Exec("DROP TABLE " + tbName)

	if err != nil {
		errMsg := strings.Split(err.Error(), ": ")
		if strings.HasPrefix(errMsg[len(errMsg)-1], "Unknown table") {
			return fmt.Errorf("unknown table")
		}
		return err
	}

	n, err := result.RowsAffected()
	if n != 0 {
		return fmt.Errorf("DROP ERROR")
	}

	return err
}

func InitTable(db *sql.DB, dbName string, tbName string) error {

	if !IsExistTable(db, dbName, tbName) {
		return fmt.Errorf("table is not exist")
	}

	err := DropTable(db, tbName)
	LogError(err)
	err = CreateTable(db, tbName)
	return err
}

func Insert(db *sql.DB, qTemplate string, params ...interface{}) error {

	result, err := db.Exec(qTemplate, params...)
	LogError(err)

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted")
	}
	return err
}

func Select(db *sql.DB, qTemplate string, param string) (interface{}, error) {

	var value interface{}
	err := db.QueryRow(qTemplate, param).Scan(&value)
	return value, err
}

func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
