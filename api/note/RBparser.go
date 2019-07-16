package note

import (
	"encoding/json"
	"io"
	"nottu/api"
)

type noteRequestBody struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

func parseNoteRequestBody(body io.ReadCloser) (*noteRequestBody, error) {
	nrb := new(noteRequestBody)
	err := json.NewDecoder(body).Decode(nrb)
	if err != nil {
		return nil, api.ErrBadRequest
	}
	return nrb, nil
}
