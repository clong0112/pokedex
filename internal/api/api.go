package api

import (
	"net/http"
	"time"

	"github.com/clong0112/pokedex/internal/cache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct{
	httpClient http.Client
	cache cache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client {
		cache: cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}


