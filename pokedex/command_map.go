package main

import (
	"fmt"
	"net/http"
)

func commandMap(cfg *config) error {
	// choose url
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	}

	// request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//temp
	fmt.Println("status:", resp.Status)

	// decode JSON

	return nil
}

func commandMapB(cfg *config) error {
	return nil
}
