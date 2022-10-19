package models

import "encoding/json"

type Episodio struct {
	FormName     string        `json:"form_name"`
	FormNames    []interface{} `json:"form_names"`
	FormOrder    int64         `json:"form_order"`
	ID           int64         `json:"id"`
	IsBattleOnly bool          `json:"is_battle_only"`
	IsDefault    bool          `json:"is_default"`
	IsMega       bool          `json:"is_mega"`
	Name         string        `json:"name"`
	Names        []interface{} `json:"names"`
	Order        int64         `json:"order"`
	Pokemon      Pokemon       `json:"pokemon"`
	Sprites      Sprites       `json:"sprites"`
	Types        []Type        `json:"types"`
	VersionGroup Pokemon       `json:"version_group"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Sprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type Type struct {
	Slot int64   `json:"slot"`
	Type Pokemon `json:"type"`
}

func UnmarshalEpisodio(data []byte) (Episodio, error) {
	var r Episodio
	err := json.Unmarshal(data, &r)
	return r, err
}
