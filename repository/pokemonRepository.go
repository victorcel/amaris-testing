package repository

import "github.com/victorcel/amaris-testing/models"

type PokemonRepository interface {
	GetPokemon(pokemonID int) (episodio models.Episodio, err error)
}

var implementationPokemon PokemonRepository

func SetUserRepository(repository PokemonRepository) {
	implementationPokemon = repository
}

func GetPokemon(pokemonID int) (episodio models.Episodio, err error) {
	return implementationPokemon.GetPokemon(pokemonID)
}
