package server

import (
	"net/http"
	"nottu/api/note"
	"nottu/app"
)

func Run() {
	nottu := app.New()
	nottu.Handle("/notes", note.EndPoint).Methods(http.MethodGet, http.MethodPost)
	nottu.Handle("/notes/{noteID}", note.ByIDEndPoint).Methods(http.MethodGet, http.MethodDelete, http.MethodPut)
	nottu.Run()
}
