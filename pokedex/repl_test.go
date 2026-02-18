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

func TestCleanInput(t *testing.T) {
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
		"you must provide a location name")

	// Test case: too many arguments
	err = commandExplore(cfg, "pastoria", "extra-arg")
	assertEqual(t, "explore with too many arguments",
		err.Error(),
		"you must provide a location name")
}

type mockPokeAPI struct {
	mockLocationResp pokeapi.Location
	mockListResp     pokeapi.RespShallowLocations
	mockError        error
	mockPokemonResp map[string]pokeapi.Pokemon
}

func (m *mockPokeAPI) GetPokemon(name string) (pokeapi.Pokemon, error) {
	if m.mockError != nil {
		return pokeapi.Pokemon{}, m.mockError
	}
	return m.mockPokemonResp[name], nil
}

func (m *mockPokeAPI) GetLocation(name string) (pokeapi.Location, error) {
	return m.mockLocationResp, m.mockError
}

func (m *mockPokeAPI) ListLocations(url *string) (pokeapi.RespShallowLocations, error) {
	return m.mockListResp, m.mockError
}

func TestCommandMap_UpdatesConfig(t *testing.T) {
	next := pokeapi.BaseURL + "/location-area/?offset=20&limit=20"
	prev := pokeapi.BaseURL + "/location-area/?offset=0&limit=20"

	mockClient := &mockPokeAPI{
		mockListResp: pokeapi.RespShallowLocations{
			Next:     &next,
			Previous: &prev,
			Results: []pokeapi.ShallowLocation{
				{Name: "location-1"},
			},
		},
	}
	//mockPokemon := &mockPokemon

	cfg := &config{
		pokeAPIClient: mockClient,
	}

	err := commandMap(cfg)
	assertEqual(t, "map should not error", err, nil)

	assertEqual(t, "config.nestsLocationURL should be updated", cfg.nextLocationsURL, &next)
	assertEqual(t, "config.prevLocationURL should be updated", cfg.prevLocationsURL, &prev)
}

func TestCommandExplore_Success(t *testing.T) {
	expectedLocation := pokeapi.Location{
		Name: "pastoria-city-area",
		PokemonEncounters: []pokeapi.PokemonEncounters{
			{Pokemon: pokeapi.NamedAPIResource{Name: "magikarp"}},
			{Pokemon: pokeapi.NamedAPIResource{Name: "gyarados"}},
		},
	}

	cfg := &config{
		pokeAPIClient: &mockPokeAPI{
			mockLocationResp: expectedLocation,
		},
	}

	err := commandExplore(cfg, "pastora")

	assertEqual(t, "should not error on valid location", err, nil)
}

func TestCatch(t *testing.T) {
	
	cfg := &config{
		pokeAPIClient: &mockPokeAPI{
		},
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	err := commandCatch(cfg, "magikarp")
	assertEqual(t, "should not error on pokemon", err, nil)
}
