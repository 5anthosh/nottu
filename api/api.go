package api

import (
	"errors"

	"github.com/5anthosh/nottu/db/note"
)

type ResponseBuilder struct {
	JSONData   *Response      `json:"data,omitempty"`
	JSONErrors *ErrorMessages `json:"errors,omitempty"`
}

//Data sets data of json
func (r *ResponseBuilder) Data() *Response {
	if r.JSONData == nil {
		r.JSONData = new(Response)
	}
	return r.JSONData
}
func (r *ResponseBuilder) Error() *ErrorMessages {
	if r.JSONErrors == nil {
		r.JSONErrors = new(ErrorMessages)
	}
	return r.JSONErrors
}

//Response containes requested info if request is successfull
type Response struct {
	JSONID      string       `json:"id,omitempty"`
	JSONMessage string       `json:"message,omitempty"`
	JSONNote    *note.Note   `json:"note,omitempty"`
	JSONNotes   []*note.Note `json:"notes,omitempty"`
}

//ID set
func (d *Response) ID(id string) *Response {
	d.JSONID = id
	return d
}

//Message #
func (d *Response) Message(message string) *Response {
	d.JSONMessage = message
	return d
}

func (d *Response) Note(note *note.Note) *Response {
	d.JSONNote = note
	return d
}
func (d *Response) Notes(notes []*note.Note) *Response {
	d.JSONNotes = notes
	return d
}

//ErrorMessages contains error
type ErrorMessages struct {
	JSONErrorMessage string `json:"message,omitempty"`
}

//Message #
func (e *ErrorMessages) Message(message string) *ErrorMessages {
	e.JSONErrorMessage = message
	return e
}

//BadRequest invalid request
const (
	BadRequest = "Bad request"
)

//ErrBadRequest invalid request error
var (
	ErrBadRequest = errors.New(BadRequest)
)
