package note

import (
	"net/http"
	"nottu/api"
	"nottu/app"
	"nottu/db/note"
	"nottu/db/status"
)

func create(c *app.Context, code *int, response *api.ResponseBuilder) {
	requestBody, err := parseNoteRequestBody(c.Request.Body)
	if err != nil {
		response.Error().Message(err.Error())
		*code = status.BadRequest
	} else {
		var id, result string
		id, result, *code, err = note.Create(c.DB, &requestBody.Title, &requestBody.Content)
		c.AppendError(err)
		if *code == status.CreatedSuccess {
			response.Data().ID(id).Message(result)
		} else {
			response.Error().Message(result)
		}
	}
}

//EndPoint #
func EndPoint(c *app.Context) {
	code := new(int)
	response := new(api.ResponseBuilder)
	switch c.Request.Method {
	case http.MethodPost:
		create(c, code, response)
	default:
		c.HTTPStatusCode(status.NotFound)
		return
	}
	c.HTTPStatusCode(*code)
	c.JSON(response)
}
