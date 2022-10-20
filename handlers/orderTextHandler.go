package handlers

import (
	"encoding/json"
	"github.com/victorcel/amaris-testing/helper"
	"github.com/victorcel/amaris-testing/useCases"
	"net/http"
)

func OrderTextHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var requestData map[string]string

		writer.Header().Set("Content-Type", "application/json")

		err := json.NewDecoder(request.Body).Decode(&requestData)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		orderText, err := useCases.OrderText(requestData["data"])

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			helper.ErrorResponse(writer, err)
			return
		}

		helper.SuccessOrderTextResponse(writer, orderText)
	}

}
