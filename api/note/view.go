package note

import (
	"net/http"

	"github.com/5anthosh/mint"
)

//New creates HandlersGroup for note
func New() *mint.HandlersGroup {
	return mint.NewGroup("/notes").
		Handler(
			mint.HandlerBuilder().
				Handle(EndPoint).
				Methods(http.MethodGet, http.MethodPost).
				Compressed(true),
		).
		Handler(
			mint.HandlerBuilder().
				Path("/"+mint.URLVar(noteIDURLVar)).
				Handle(ByIDEndPoint).
				Methods(http.MethodGet, http.MethodDelete, http.MethodPut).
				Compressed(true),
		)
}
