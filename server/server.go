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
	"github.com/gorilla/handlers"
)

const defaultPort = "1313"
const defaultDBFileLocation = ".config/nottu/"
const databaseName = ".nottu.sqlite3"

//Run starts server
func Run() {
	nottu := mint.New()
	nottu.Group("/notes").
		Handler(
			new(mint.HandlersContext).
				Path("").
				Handle(note.EndPoint).
				Methods(http.MethodGet, http.MethodPost).
				Compressed(true),
		).Handler(
		new(mint.HandlersContext).
			Path("/{"+note.NoteIDURLVar+"}").
			Handle(note.ByIDEndPoint).
			Methods(http.MethodGet, http.MethodDelete, http.MethodPut).Compressed(true),
	)
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

	router := nottu.Build()
	port := defaultPort
	serverAdd := ":" + port
	fmt.Println("🚀  Starting server....")

	//	go open(localAddress + "/#/pltm/container")
	protocal := "http"
	localAddress := protocal + "://localhost" + serverAdd
	fmt.Println("🌠 Ready on " + localAddress)
	err = http.ListenAndServe(serverAdd, handlers.RecoveryHandler()(router))
	if err != nil {
		fmt.Println("Stopping the server" + err.Error())
	}

}
