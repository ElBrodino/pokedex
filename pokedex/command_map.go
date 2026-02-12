package main

import (
	"fmt"
)

func commandMap(cfg *config) error {

	// request
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
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

func commandMapB(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
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
