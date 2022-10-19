package useCases

import (
	"errors"
	"fmt"
	"github.com/victorcel/amaris-testing/models"
	"io"
	"net/http"
)

const (
	NOEXISTPOKEMONID = "no exist Pokemon ID"
)

type Pokemon struct {
	client *http.Client
}

func NewPokemon(client *http.Client) *Pokemon {
	return &Pokemon{client: client}
}

func (p *Pokemon) GetPokemon(pokemonID int) (models.Episodio, error) {

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%d", pokemonID)
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	episodio := models.Episodio{}

	if err != nil {

		return episodio, err
	}

	res, err := p.client.Do(req)

	if err != nil {

		return episodio, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return episodio, err
	}

	episodio, err = models.UnmarshalEpisodio(body)

	if err != nil {
		return episodio, errors.New(NOEXISTPOKEMONID)
	}

	return episodio, nil
}
