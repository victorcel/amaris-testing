package main

import (
	"fmt"
	"github.com/victorcel/amaris-testing/useCases"
	"net/http"
)

func main() {

	//fmt.Println(useCases.OrderText("Luis,Camilo;Andres,Laura"))
	usePokemon := useCases.NewPokemon(&http.Client{})
	pokemon, err := usePokemon.GetPokemon(1111)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pokemon.ID)
}
