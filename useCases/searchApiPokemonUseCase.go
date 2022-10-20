package useCases

import (
	"errors"
	"fmt"
	"github.com/victorcel/amaris-testing/models"
	"io"
	"net/http"
	"os"
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

func (p *Pokemon) GetPokemon(pokemonID int) (episodio models.Episodio, err error) {

	url := fmt.Sprintf("%s%d", os.Getenv("POKEMON_API"), pokemonID)
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

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
