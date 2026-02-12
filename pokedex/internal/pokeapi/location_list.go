package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type ReapShallowlocations struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (ReapShallowlocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, ok := c.cache.Get(url)
	if ok {
		locationResp := ReapShallowlocations{}
		json.Unmarshal(val, &locationResp)
		return locationResp, nil
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
