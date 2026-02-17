package main

import "pokedex/internal/pokeapi"

type pokeAPIClient interface {
	GetLocation(locationName string) (pokeapi.Location, error)
	ListLocations(pageURL *string) (pokeapi.RespShallowLocations, error)
	GetPokemon(name string) (pokeapi.Pokemon, error)
}

type config struct {
	pokeAPIClient    pokeAPIClient //interface
	nextLocationsURL *string
	prevLocationsURL *string

	caughtPokemon map[string]pokeapi.Pokemon
}
