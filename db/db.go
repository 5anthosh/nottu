package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/rs/xid"
)

const (
	sqliteEngine = "sqlite3"
)

//Db modes
const (
	DevMode  = "Dev"
	TestMode = "Test"
)

func sqliteDBConnection(devdb string, testdb string, mode string) *sql.DB {
	var DatabaseName string
	if mode == TestMode {
		DatabaseName = testdb
	} else {
		DatabaseName = devdb
	}
	if _, err := os.Stat(DatabaseName); os.IsNotExist(err) {
		log.Fatal(DatabaseName + " does not exist : reinit database to create")
	}
	database, err := sql.Open(sqliteEngine, DatabaseName)
	if err != nil {
		log.Fatal("error in opening database connection")
	}
	return database
}

func sqliteWithOutDBConnection(devdb string, testdb string, mode string) *sql.DB {
	var DatabaseName string
	if mode == TestMode {
		DatabaseName = testdb
	} else {
		DatabaseName = devdb
	}
	os.Remove(DatabaseName)
	f, err := os.Create(DatabaseName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	f.Close()
	return sqliteDBConnection(devdb, testdb, mode)
}

//ReinitDB reinits the database , it is dangerous function because it wipes out every data
func ReinitDB() {

}

//GenerateUniqueID generates uuid
func GenerateUniqueID() string {
	return xid.New().String()
}
