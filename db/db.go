package db

import (
	"database/sql"
	"log"
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

type Database struct {
	DevDatabaseFilePath string
}

func (db Database) Connection() *sql.DB {
	if _, err := os.Stat(db.DevDatabaseFilePath); os.IsNotExist(err) {
		return db.withOutDBConnection()
	}
	database, err := sql.Open(sqliteEngine, db.DevDatabaseFilePath)
	if err != nil {
		log.Fatal("error in opening database connection", err)
	}
	return database
}

func (db Database) withOutDBConnection() *sql.DB {
	f, err := os.Create(db.DevDatabaseFilePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	f.Close()
	return db.Connection()
}

// //ReinitDB reinits the database , it is dangerous function because it wipes out every data
// func ReinitDB() {
// 	os.Remove(config.DatabaseName)
// 	db := sqliteWithOutDBConnection(config.DatabaseName, "test.sqite3", DevMode)
// 	tx, err := db.Begin()
// 	if err == nil {
// 		_, err = tx.Exec(noteQuery)
// 		if err == nil {
// 			tx.Commit()
// 			return
// 		}
// 	}
// 	tx.Rollback()
// 	log.Println("Error in creating database")
// }

//GenerateUniqueID generates uuid
func GenerateUniqueID() string {
	return xid.New().String()
}
