package pokeapi

type RespShallowCreatures struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type ShallowLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RespShallowLocations struct {
	Next     *string           `json:"next"`
	Previous *string           `json:"previous"`
	Results  []ShallowLocation `json:"results"`
}

type Location struct {
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name:"`
	URL  string `json:"url"`
}
