package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide one a location name")
	}
	location, err := cfg.pokeAPIClient.GetLocation(args[0])
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	fmt.Printf("Exploring %v...\n", location.Name)
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
		//fmt.Printf("%v", location.Name)
	}
	//fmt.Print(location)
	return nil
}
