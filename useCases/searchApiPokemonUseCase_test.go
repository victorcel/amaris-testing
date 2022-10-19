package useCases

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/victorcel/amaris-testing/models"
	"net/http"
	"reflect"
	"testing"
)

func TestPokemon_GetPokemon(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%d", 1),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, models.Episodio{
				ID:   1,
				Name: "Victor",
			})
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%d", 1111),
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusInternalServerError, ""), nil
		})

	httpmock.RegisterResponder("GET1", fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%d", 22),
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusInternalServerError, ""), nil
		})

	type fields struct {
		client *http.Client
	}

	type args struct {
		pokemonID int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Episodio
		mocker  func(a args)
		wantErr bool
	}{
		{
			name: "TestGetPokemon success",
			fields: fields{
				client: &http.Client{},
			},
			args: args{
				pokemonID: 1,
			},
			want: *&models.Episodio{
				ID:   1,
				Name: "Victor",
			},
			wantErr: false,
		},
		{
			name: "TestGetPokemon error",
			fields: fields{
				client: &http.Client{},
			},
			args: args{
				pokemonID: 1111,
			},
			want: *&models.Episodio{},

			wantErr: true,
		},
		{
			name: "TestGetPokemon error methodErr",
			fields: fields{
				client: &http.Client{},
			},
			args: args{
				pokemonID: 22,
			},
			want: *&models.Episodio{},

			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			p := &Pokemon{
				client: tt.fields.client,
			}
			got, err := p.GetPokemon(tt.args.pokemonID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPokemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPokemon() got = %v, want %v", got, tt.want)
			}
		})
	}
}
