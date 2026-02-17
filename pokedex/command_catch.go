package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide one a pokemon name")
	}

	name := strings.ToLower(args[0])
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeAPIClient.GetPokemon(name)
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)

	if chance < 10 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}
	fmt.Printf("%s was caught!\n", name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
