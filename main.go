package main

import (
	"github.com/joho/godotenv"
	"github.com/victorcel/amaris-testing/repository"
	"github.com/victorcel/amaris-testing/routers"
	"github.com/victorcel/amaris-testing/server"
	"github.com/victorcel/amaris-testing/useCases"
	"log"
	"net/http"
	"os"
)

func main() {

	//fmt.Println(useCases.OrderText("Luis,Camilo;Andres,Laura"))
	//usePokemon := useCases.NewPokemon(&http.Client{})
	//pokemon, err := usePokemon.GetPokemon(1111)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(pokemon.ID)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s, err := server.NewServer(&server.Config{
		Port: os.Getenv("PORT"),
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}
	pokemon := useCases.NewPokemon(&http.Client{})
	repository.SetUserRepository(pokemon)
	s.Start(routers.BindRoutes)
}
