package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type ReapShallowCreatures struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type ReapShallowlocations struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
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

func (c *Client) ListLocations(pageURL *string) (ReapShallowlocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, ok := c.cache.Get(url)
	if ok {
		locationReap := ReapShallowlocations{}
		json.Unmarshal(val, &locationReap)
		return locationReap, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ReapShallowlocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ReapShallowlocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ReapShallowlocations{}, err
	}

	var out ReapShallowlocations
	if err := json.Unmarshal(dat, &out); err != nil {
		return ReapShallowlocations{}, err
	}

	c.cache.Add(url, dat)
	return out, nil
}
