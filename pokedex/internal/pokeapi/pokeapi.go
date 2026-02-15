package pokeapi

import (
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache:      *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{Timeout: timeout},
	}
}
