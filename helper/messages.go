package helper

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error   error
	Message interface{}
	Code    int
}

func ErrorResponse(writer http.ResponseWriter, err error) {
	json.NewEncoder(writer).Encode(Error{
		Error:   err,
		Message: err.Error(),
		Code:    http.StatusBadRequest,
	})
}

func SuccessResponse(writer http.ResponseWriter, data interface{}) {
	json.NewEncoder(writer).Encode(data)
}
