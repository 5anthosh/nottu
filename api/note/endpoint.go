package note

import (
	"database/sql"
	"net/http"

	"github.com/5anthosh/nottu/api"
	"github.com/5anthosh/nottu/core/note"
	"github.com/5anthosh/nottu/core/status"

	"github.com/5anthosh/mint"
)

const noteIDURLVar = "noteID"

func create(r *http.Request, DB *sql.DB, response *api.ResponseBuilder) (err error) {
	requestBody, err := parseCreateRequestBody(r.Body)
	if err != nil {
		response.Code(http.StatusBadRequest).Error().Message(err.Error())
		return
	}
	id, result, code, err := note.Create(DB, requestBody.Title, requestBody.Content)
	response.Code(code)
	if code == status.CreatedSuccess {
		response.Data().ID(id).Message(result)
		return
	}
	response.Error().Message(result)
	return
}

func get(DB *sql.DB, response *api.ResponseBuilder) (err error) {
	notes, result, code, err := note.Get(DB)
	response.Code(code)
	if code == status.OK {
		response.Data().Notes(notes)
		return
	}
	response.Error().Message(result)
	return
}

func byID(DB *sql.DB, id string, response *api.ResponseBuilder) (err error) {
	noteObj, result, code, err := note.ByID(DB, id)
	response.Code(code)
	if code == status.OK {
		response.Data().Note(noteObj)
		return
	}
	response.Error().Message(result)
	return
}

func delete(DB *sql.DB, id string, response *api.ResponseBuilder) (err error) {
	_, code, err := note.Delete(DB, id)
	response.Code(code)
	return
}

func update(r *http.Request, DB *sql.DB, id string, response *api.ResponseBuilder) (err error) {
	requestBody, err := parseUpdateRequestBody(r.Body)
	if err == nil {
		var result string
		var code int
		result, code, err = note.Update(DB, id, requestBody.Title, requestBody.Content)
		response.Code(code)
		if code == status.OK {
			response.Data().Message(result)
			return
		}
		response.Error().Message(result)
		return

	}
	response.Code(status.BadRequest).Error().Message(err.Error())
	return
}

//EndPoint #
func EndPoint(c *mint.Context) {
	response := new(api.ResponseBuilder)
	value, _ := c.MintParam("DB")
	DB := value.(*sql.DB)
	var err error
	switch c.Req.Method {
	case http.MethodGet:
		err = get(DB, response)
		break
	case http.MethodPost:
		err = create(c.Req, DB, response)
		break
	default:
		c.Status(status.NotFound)
		return
	}
	c.Error(err)
	c.JSON(response.HTTPCode, response)
}

//ByIDEndPoint #
func ByIDEndPoint(c *mint.Context) {
	response := new(api.ResponseBuilder)
	noteID, _ := c.Param(noteIDURLVar)
	value, _ := c.MintParam("DB")
	DB := value.(*sql.DB)
	var err error
	switch c.Req.Method {
	case http.MethodGet:
		err = byID(DB, noteID, response)
		break
	case http.MethodPut:
		err = update(c.Req, DB, noteID, response)
		break
	case http.MethodDelete:
		err = delete(DB, noteID, response)
		break
	default:
		c.Status(status.NotFound)
		return
	}
	c.Error(err)
	c.JSON(response.HTTPCode, response)
}
