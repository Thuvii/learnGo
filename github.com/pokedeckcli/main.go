package main

import (
	"time"

	"github.com/Thuvii/pokedeckcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := &config{
		pokeapiClient:  client,
		pokemonCatched: make(map[string]pokeapi.Pokemon),
	}
	startRepl(config)
}
