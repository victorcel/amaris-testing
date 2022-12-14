package helper

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error   error       `json:"error"`
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
}

type successResponse struct {
	Data interface{} `json:"data" json:"data,omitempty"`
}

type successOrderTextResponse struct {
	Quantity int         `json:"quantity"`
	Data     interface{} `json:"data" json:"data,omitempty"`
}

func ErrorResponse(writer http.ResponseWriter, err error) {
	json.NewEncoder(writer).Encode(errorResponse{
		Error:   err,
		Message: err.Error(),
		Code:    http.StatusBadRequest,
	})
}

func SuccessResponse(writer http.ResponseWriter, data interface{}) {
	json.NewEncoder(writer).Encode(successResponse{
		Data: data,
	})
}
func SuccessOrderTextResponse(writer http.ResponseWriter, data []string) {
	json.NewEncoder(writer).Encode(successOrderTextResponse{
		Quantity: len(data),
		Data:     data,
	})
}
