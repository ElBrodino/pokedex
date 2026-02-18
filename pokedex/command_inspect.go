package main

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide one a pokemon name")
	}

	name := strings.ToLower(args[0])
	creature := cfg.caughtPokemon[name]
	fmt.Printf("Name: %s\n", creature.Name)
	fmt.Printf("Height: %d\n", creature.Height)
	fmt.Printf("Weight: %d\n", creature.Weight)
	fmt.Println("Stats:")
	for _, x := range creature.Stats {
		fmt.Printf("-%s: %d\n", x.Stat.Name, x.BaseStat)
	}
	fmt.Println("Types:")
	for _, x := range creature.Types {
		fmt.Printf("- %s", x.Type.Name)
	}

	return nil
}
