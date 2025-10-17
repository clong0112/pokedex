package main

import (
	"time"

	"github.com/clong0112/pokedex/internal/api"
)

type config struct {
	apiClient               api.Client
	NextLocationAreaURL     *string
	PreviousLocationAreaURL *string
	caughtPokemon map[string]api.Pokemon
}

func main() {
	cfg := config{
		apiClient: api.NewClient(time.Hour),
		caughtPokemon: make(map[string]api.Pokemon),
	}

	startRepl(&cfg)
}
