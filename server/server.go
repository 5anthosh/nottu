package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/user"
	"path/filepath"

	"github.com/5anthosh/nottu/api/note"
	"github.com/5anthosh/nottu/core/database"

	"github.com/5anthosh/mint"
	"github.com/gorilla/mux"
)

const defaultPort = "1313"
const defaultDBFileLocation = ".config/nottu/"
const databaseName = "/.nottu.sqlite3"
const testDataBaseName = "/.nottu_test.sqlite3"

//Run starts server
func Run() {
	router, DB := Build()
	defer DB.Close()
	port := defaultPort
	serverAdd := ":" + port
	fmt.Println("🚀  Starting server....")

	//	go open(localAddress + "/#/pltm/container")
	protocal := "http"
	localAddress := protocal + "://localhost" + serverAdd
	fmt.Println("🌠 Ready on " + localAddress)
	err := http.ListenAndServe(serverAdd, router)
	if err != nil {
		fmt.Println("Stopping the server" + err.Error())
	}

}

//Build builds router with views config
func Build() (*mux.Router, *sql.DB) {
	nottu := mint.New()
	nottu.AddGroup(note.New())
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fileLocation := filepath.Join(usr.HomeDir, defaultDBFileLocation)
	database := database.Database{
		FileParentPath:      fileLocation,
		DevDatabaseFilePath: fileLocation + databaseName,
	}
	DB := database.Connection()
	nottu.Set("DB", DB)
	return nottu.Build(), DB
}

func TestBuild() (*mux.Router, *sql.DB) {
	nottu := mint.Simple()
	nottu.AddGroup(note.New())
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	usr, err = user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fileLocation := filepath.Join(usr.HomeDir, defaultDBFileLocation)
	database := database.Database{
		FileParentPath:      fileLocation,
		DevDatabaseFilePath: fileLocation + testDataBaseName,
	}
	DB := database.Connection()
	nottu.Set("DB", DB)
	return nottu.Build(), DB
}
