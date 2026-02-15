package pokeapi

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name:"`
	URL  string `json:"url"`
}
