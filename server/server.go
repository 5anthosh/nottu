package server

import (
	"fmt"
	"log"
	"net/http"
	"os/user"
	"path/filepath"

	"github.com/5anthosh/nottu/api/note"
	"github.com/5anthosh/nottu/db"

	"github.com/5anthosh/mint"
	"github.com/gorilla/mux"
)

const defaultPort = "1313"
const defaultDBFileLocation = ".config/nottu/"
const databaseName = "/.nottu.sqlite3"

//Run starts server
func Run() {
	router := Build()
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
func Build() *mux.Router {
	nottu := mint.New()
	nottu.AddGroup(note.New())
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fileLocation := filepath.Join(usr.HomeDir, defaultDBFileLocation)
	nottu.RegisterDB(
		db.Database{
			FileParentPath:      fileLocation,
			DevDatabaseFilePath: fileLocation + databaseName,
		},
	)
	return nottu.Build()
}

func TestBuild() *mux.Router {
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
	nottu.RegisterDB(
		db.Database{
			FileParentPath:      fileLocation,
			DevDatabaseFilePath: fileLocation + "/.nottu_test.sqlite3",
		},
	)
	return nottu.Build()
}
