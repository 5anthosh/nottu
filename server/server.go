package server

import (
	"net/http"
	"nottu/api/note"
	"nottu/app"
)

func Run() {
	nottu := app.New()
	nottu.Handle("/notes", note.EndPoint).Methods(http.MethodPost)
	nottu.Run()
}
