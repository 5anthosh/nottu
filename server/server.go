package server

import (
	"fmt"
	"net/http"
	"nottu/api/note"
	"nottu/config"
	"nottu/db"

	"github.com/5anthosh/mint"
	"github.com/gorilla/handlers"
)

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
	nottu.RegisterDB(
		db.Database{
			DevDatabaseFilePath: config.DatabaseName,
		},
	)
	router := nottu.Build()
	port := config.DefaultPort
	serverAdd := ":" + port
	fmt.Println("🚀  Starting server....")

	//	go open(localAddress + "/#/pltm/container")
	protocal := "http"
	var err error
	localAddress := protocal + "://localhost" + serverAdd
	fmt.Println("🌠 Ready on " + localAddress)
	err = http.ListenAndServe(serverAdd, handlers.RecoveryHandler()(router))
	if err != nil {
		fmt.Println("Stopping the server" + err.Error())
	}

}
