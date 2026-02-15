package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {

	// request
	reap, err := cfg.pokeAPIClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	for _, r := range reap.Results {
		fmt.Println(r.Name)
	}
	cfg.nextLocationsURL = reap.Next
	cfg.prevLocationsURL = reap.Previous

	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := cfg.pokeAPIClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}
	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	return nil
}
