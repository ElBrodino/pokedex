package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := BaseURL + "/pokemon/" + pokemonName
	out := Pokemon{}

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &out); err != nil {
			return Pokemon{}, err
		}
		return out, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("pokemon not found: %s", pokemonName)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if err := json.Unmarshal(dat, &out); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return out, nil
}
