package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Second)

	cfg := &config{
		pokeAPIClient: client,
	}
	startRepl(cfg)
}
