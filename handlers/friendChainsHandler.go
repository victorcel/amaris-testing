package handlers

import (
	"encoding/json"
	"errors"
	"github.com/victorcel/amaris-testing/helper"
	"github.com/victorcel/amaris-testing/useCases"
	"net/http"
)

func FriendChainsHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var requestData struct {
			X string `json:"x"`
			Y string `json:"y"`
		}

		writer.Header().Set("Content-Type", "application/json")

		err := json.NewDecoder(request.Body).Decode(&requestData)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			helper.ErrorResponse(writer, err)
			return
		}

		dataX := requestData.X
		dataY := requestData.Y

		if dataX == "" || dataY == "" {
			writer.WriteHeader(http.StatusBadRequest)
			helper.ErrorResponse(writer, errors.New("empty variables X or Y"))
			return
		}

		orderText, err := useCases.FriendlyChains(dataX, dataY)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			helper.ErrorResponse(writer, err)
			return
		}

		helper.SuccessResponse(writer, orderText)
	}

}
