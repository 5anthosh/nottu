package status

import "net/http"

//http statuscode
const (
	CreatedSuccess = http.StatusCreated             //201
	OK             = http.StatusOK                  //200
	NotFound       = http.StatusNotFound            //404
	NotAuthorized  = http.StatusUnauthorized        //401
	AlreadyExists  = http.StatusConflict            //409
	ServerError    = http.StatusInternalServerError //500
	BadRequest     = http.StatusBadRequest          //400
	DeletedSuccess = http.StatusNoContent           //204
)
