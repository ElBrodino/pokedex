package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type RespShallowCreatures struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type ShallowLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RespShallowLocations struct {
	Next     *string           `json:"next"`
	Previous *string           `json:"previous"`
	Results  []ShallowLocation `json:"results"`
}

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	out := Location{}

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &out); err != nil {
			return Location{}, err
		}
		return out, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	//err = json.Unmarshal(dat, &out)
	if err := json.Unmarshal(dat, &out); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	return out, nil

}

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, ok := c.cache.Get(url)
	if ok {
		locationReap := RespShallowLocations{}
		json.Unmarshal(val, &locationReap)
		return locationReap, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	var out RespShallowLocations
	if err := json.Unmarshal(dat, &out); err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)
	return out, nil
}
