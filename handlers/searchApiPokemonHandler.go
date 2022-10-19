package handlers

import (
	"github.com/gorilla/mux"
	"github.com/victorcel/amaris-testing/helper"
	"github.com/victorcel/amaris-testing/repository"
	"net/http"
	"strconv"
)

func SearchPokemonHandler() http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		params := mux.Vars(request)

		var id, _ = strconv.Atoi(params["id"])
		pokemon, err := repository.GetPokemon(id)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			helper.ErrorResponse(writer, err)
			return
		}

		helper.SuccessResponse(writer, pokemon)
	}
}
