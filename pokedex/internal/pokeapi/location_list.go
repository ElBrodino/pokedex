package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := BaseURL + "/location-area/" + locationName
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

	if resp.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("location not found: %s", locationName)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	if err := json.Unmarshal(dat, &out); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	return out, nil

}

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := BaseURL + "/location-area"
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

	if resp.StatusCode != http.StatusOK {
		return RespShallowLocations{}, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

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
