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
	for i := range cfg.caughtPokemon {
		if i == name {
			fmt.Printf("")
		}
	}
	return nil
}
