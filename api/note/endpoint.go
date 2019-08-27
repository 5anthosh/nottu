package note

import (
	"net/http"

	"github.com/5anthosh/nottu/api"
	"github.com/5anthosh/nottu/db/note"
	"github.com/5anthosh/nottu/db/status"

	"github.com/5anthosh/mint"
)

const NoteIDURLVar = "noteID"

func create(c *mint.Context, code *int, response *api.ResponseBuilder) {
	requestBody, err := parseCreateRequestBody(c.Request.Body)
	if err != nil {
		response.Error().Message(err.Error())
		*code = status.BadRequest
	} else {
		var id, result string
		id, result, *code, err = note.Create(c.DB, requestBody.Title, requestBody.Content)
		c.AppendError(err)
		if *code == status.CreatedSuccess {
			response.Data().ID(id).Message(result)
		} else {
			response.Error().Message(result)
		}
	}
}
func get(c *mint.Context, code *int, response *api.ResponseBuilder) {
	var result string
	var notes []note.Note
	var err error
	notes, result, *code, err = note.Get(c.DB)
	c.AppendError(err)
	if *code != status.OK {
		response.Error().Message(result)
	} else {
		response.Data().Notes(notes)
	}
}

func byID(c *mint.Context, id string, code *int, response *api.ResponseBuilder) {
	var result string
	var noteObj note.Note
	var err error
	noteObj, result, *code, err = note.ByID(c.DB, id)
	c.AppendError(err)
	if *code != status.OK {
		response.Error().Message(result)
	} else {
		response.Data().Note(noteObj)
	}
}

func delete(c *mint.Context, id string, code *int, response *api.ResponseBuilder) {
	var err error
	var result string
	result, *code, err = note.Delete(c.DB, id)
	c.AppendError(err)
	if *code != status.DeletedSuccess {
		response.Error().Message(result)
	}
}

func update(c *mint.Context, id string, code *int, response *api.ResponseBuilder) {
	requestBody, err := parseUpdateRequestBody(c.Request.Body)
	if err != nil {
		response.Error().Message(err.Error())
		*code = status.BadRequest
	} else {
		var result string
		result, *code, err = note.Update(c.DB, id, requestBody.Title, requestBody.Content)
		c.AppendError(err)
		if *code != status.OK {
			response.Error().Message(result)
		} else {
			response.Data().Message(result)
		}

	}
}

//EndPoint #
func EndPoint(c *mint.Context) {
	code := new(int)
	response := new(api.ResponseBuilder)
	switch c.Request.Method {
	case http.MethodGet:
		get(c, code, response)
	case http.MethodPost:
		create(c, code, response)
	default:
		c.Status(status.NotFound)
		return
	}
	c.JSON(*code, response)
}

//ByIDEndPoint #
func ByIDEndPoint(c *mint.Context) {
	code := new(int)
	response := new(api.ResponseBuilder)
	noteID := c.URLParams["noteID"]
	switch c.Request.Method {
	case http.MethodGet:
		byID(c, noteID, code, response)
	case http.MethodPut:
		update(c, noteID, code, response)
	case http.MethodDelete:
		delete(c, noteID, code, response)
	default:
		c.Status(status.NotFound)
		return
	}
	c.JSON(*code, response)
}
