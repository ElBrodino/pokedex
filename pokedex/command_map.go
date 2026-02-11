package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationAreaResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	// choose url
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	}

	// request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// decode JSON
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("pokeapi returned status %s", resp.Status)
	}

	var data locationAreaResponse
	err2 := json.NewDecoder(resp.Body).Decode(&data)
	if err2 != nil {
		return err2
	}
	for _, r := range data.Results {
		fmt.Println(r.Name)
	}
	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	return nil
}

func commandMapB(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = *cfg.prevLocationsURL
	}
	// request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// decode JSON
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("pokeapi returned status %s", resp.Status)
	}

	var data locationAreaResponse
	err2 := json.NewDecoder(resp.Body).Decode(&data)
	if err2 != nil {
		return err2
	}
	for _, r := range data.Results {
		fmt.Println(r.Name)
	}
	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	return nil
}
