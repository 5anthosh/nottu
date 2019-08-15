package note

import (
	"encoding/json"
	"io"

	"github.com/5anthosh/nottu/api"
)

type noteCreateRequestBody struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type noteUpdateRequestBody struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

func parseCreateRequestBody(body io.ReadCloser) (*noteCreateRequestBody, error) {
	nrb := new(noteCreateRequestBody)
	err := json.NewDecoder(body).Decode(nrb)
	if err != nil {
		return nil, api.ErrBadRequest
	}
	return nrb, nil
}
func parseUpdateRequestBody(body io.ReadCloser) (*noteUpdateRequestBody, error) {
	nrb := new(noteUpdateRequestBody)
	jd := json.NewDecoder(body)
	jd.DisallowUnknownFields()
	err := jd.Decode(nrb)
	if err != nil {
		return nil, err
	}
	return nrb, nil
}
