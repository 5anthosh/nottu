package db

import (
	"database/sql"
	"log"
	"nottu/config"
	"os"

	_ "github.com/mattn/go-sqlite3"
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

func SqliteDBConnection(devdb string, testdb string, mode string) *sql.DB {

	var DatabaseName string
	if mode == TestMode {
		DatabaseName = testdb
	} else {
		DatabaseName = devdb
	}
	if _, err := os.Stat(DatabaseName); os.IsNotExist(err) {
		return sqliteWithOutDBConnection(devdb, testdb, mode)
	}
	database, err := sql.Open(sqliteEngine, DatabaseName)
	if err != nil {
		log.Fatal("error in opening database connection", err)
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
	f, err := os.Create(DatabaseName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	f.Close()
	return SqliteDBConnection(devdb, testdb, mode)
}

//ReinitDB reinits the database , it is dangerous function because it wipes out every data
func ReinitDB() {
	os.Remove(config.DatabaseName)
	db := sqliteWithOutDBConnection(config.DatabaseName, "test.sqite3", DevMode)
	tx, err := db.Begin()
	if err == nil {
		_, err = tx.Exec(noteQuery)
		if err == nil {
			tx.Commit()
			return
		}
	}
	tx.Rollback()
	log.Println("Error in creating database")
}

//GenerateUniqueID generates uuid
func GenerateUniqueID() string {
	return xid.New().String()
}
