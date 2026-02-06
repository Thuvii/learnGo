package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, pokemonName ...string) error {
	if len(cfg.pokemonCatched) == 0 {
		return errors.New("You haven't caught any pokemon")
	}
	fmt.Println("Pokemon: ")
	for _, poke := range cfg.pokemonCatched {
		fmt.Printf("- %v\n", poke.Name)
	}
	return nil
}
