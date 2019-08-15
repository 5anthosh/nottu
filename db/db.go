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
	FileParentPath      string
	DevDatabaseFilePath string
}

func (db Database) Connection() *sql.DB {
	if _, err := os.Stat(db.DevDatabaseFilePath); os.IsNotExist(err) {
		log.Println("Datebase file not found ", db.DevDatabaseFilePath, " : Creating new database file....")
		dbConn := db.withOutDBConnection()
		reinitDB(dbConn)
		return dbConn
	}
	database, err := sql.Open(sqliteEngine, db.DevDatabaseFilePath)
	if err != nil {
		log.Fatal("error in opening database connection", err)
	}
	return database
}

func (db Database) withOutDBConnection() *sql.DB {
	err := os.MkdirAll(db.FileParentPath, os.ModePerm)
	if err != nil {
		log.Fatalln("Could not create folder ", db.FileParentPath, " ", err)
	}
	f, err := os.Create(db.DevDatabaseFilePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	f.Close()
	return db.Connection()
}

func reinitDB(dbConn *sql.DB) {
	tx, err := dbConn.Begin()
	if err == nil {
		_, err = tx.Exec(noteQuery)
		if err == nil {
			tx.Commit()
			return
		}
	}
	tx.Rollback()
	log.Fatalln("Error in creating table Note", err)
}

//GenerateUniqueID generates uuid
func GenerateUniqueID() string {
	return xid.New().String()
}
