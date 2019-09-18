package api

import (
	"errors"

	"github.com/5anthosh/nottu/core/note"
)

//ResponseBuilder contructs response message
type ResponseBuilder struct {
	HTTPCode   int            `json:"code,omitempty"`
	JSONData   *Response      `json:"data,omitempty"`
	JSONErrors *ErrorMessages `json:"errors,omitempty"`
}

//Data sets data of json
func (rb *ResponseBuilder) Data() *Response {
	if rb == nil {
		return nil
	}
	if rb.JSONData == nil {
		rb.JSONData = new(Response)
	}
	return rb.JSONData
}

//Code #
func (rb *ResponseBuilder) Code(value int) *ResponseBuilder {
	if rb == nil {
		return nil
	}
	rb.HTTPCode = value
	return rb
}

//Error sets error response
func (rb *ResponseBuilder) Error() *ErrorMessages {
	if rb == nil {
		return nil
	}
	if rb.JSONErrors == nil {
		rb.JSONErrors = new(ErrorMessages)
	}
	return rb.JSONErrors
}

//Response containes requested info if request is successfull
type Response struct {
	JSONID      string      `json:"id,omitempty"`
	JSONMessage string      `json:"message,omitempty"`
	JSONNote    *note.Note  `json:"note,omitempty"`
	JSONNotes   []note.Note `json:"notes,omitempty"`
}

//ID set
func (r *Response) ID(id string) *Response {
	if r == nil {
		return nil
	}
	r.JSONID = id
	return r
}

//Message #
func (r *Response) Message(message string) *Response {
	if r == nil {
		return nil
	}
	r.JSONMessage = message
	return r
}

//Note #
func (r *Response) Note(note note.Note) *Response {
	if r == nil {
		return nil
	}
	r.JSONNote = &note
	return r
}

//Notes #
func (r *Response) Notes(notes []note.Note) *Response {
	if r == nil {
		return nil
	}
	r.JSONNotes = notes
	return r
}

//ErrorMessages contains error
type ErrorMessages struct {
	JSONErrorMessage string `json:"message,omitempty"`
}

//Message #
func (e *ErrorMessages) Message(message string) *ErrorMessages {
	if e == nil {
		return nil
	}
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
