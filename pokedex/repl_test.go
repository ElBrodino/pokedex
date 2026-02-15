package main

import (
	"pokedex/internal/pokeapi"
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, desc, got, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\n*** %#v ***\n Expected %#v, got %#v", desc, want, got)
	}
}

func TestStuff(t *testing.T) {
	assertEqual(t, "Spaces front and back and middle",
		cleanInput("   	hello	world   "),
		[]string{"hello", "world"})
	assertEqual(t, "Spaces front and back and middle, amount",
		len(cleanInput("   	hello	world   ")),
		2)
	assertEqual(t, "mixed case words",
		cleanInput("Charmander Balbasaur PIKACHU"),
		[]string{"charmander", "balbasaur", "pikachu"})
	assertEqual(t, "empty input",
		cleanInput(""),
		[]string{""})
}

func TestCommandExplore(t *testing.T) {
	cfg := &config{}

	err := commandExplore(cfg)
	assertEqual(t, "explore with no argumnents",
		err.Error(),
		"you must provide one a location name")

	// Test case: too many arguments
	err = commandExplore(cfg, "pastoria", "extra-arg")
	assertEqual(t, "explore with too many arguments",
		err.Error(),
		"you must provide one a location name")
}

type mockPokeAPI struct {
	mockLocationResp pokeapi.Location
	mockListResp     pokeapi.RespShallowLocations
	mockError        error
}

func (m *mockPokeAPI) GetLocation(name string) (pokeapi.Location, error) {
	return m.mockLocationResp, m.mockError
}

func (m *mockPokeAPI) ListLocations(url *string) (pokeapi.RespShallowCreatures, error) {
	return m.mockListResp, m.mockError
}

func TestCommandExplore_Success(t *testing.T) {
	expectedLocation := pokeapi.Location{
		Name: "pastoria",
		PokemonEncounters: []struct {
			Pokemon struct {
				Name string `json:"Name"`
			} `json:"pokemon"`
		}{
			{Pokemon: struct {
				Name string `json:"Name"`
			}{Name: "pikachu"}},
			{Pokemon: struct {
				Name string `"json:"Name"`
			}{Name: "mew"}},
		},
	}

	cfg := &config{
		pokeAPIClient: &mockPokeAPI{
			mockResponse: expectedLocation,
		},
	}

	err := commandExplore(cfg, "pastora")

	assertEqual(t, "should not error on valid location", err, nil)
}
