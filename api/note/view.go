package note

import (
	"net/http"

	"github.com/5anthosh/mint"
)

//New creates HandlersGroup for note
func New() *mint.HandlersGroup {
	return mint.NewGroup("/notes").
		Handler(
			new(mint.HandlersContext).
				Handle(EndPoint).
				Methods(http.MethodGet, http.MethodPost).
				Compressed(true),
		).Handler(
		new(mint.HandlersContext).
			Path("/"+mint.URLVar(NoteIDURLVar)).
			Handle(ByIDEndPoint).
			Methods(http.MethodGet, http.MethodDelete, http.MethodPut).Compressed(true),
	)
}
