package main

import "net/http"

func commandMap(cfg *config) error {
	if cfg.nextLocationsURL == nil {
		url := "https://pokeapi.co/api/v2/location-area/"

	} else {
		url := *cfg.nextLocationsURL
	}
	http.Get(url)
	defer resp.Body.Close()
}

func commandMapB(cfg *config) error {
	return nil
}
