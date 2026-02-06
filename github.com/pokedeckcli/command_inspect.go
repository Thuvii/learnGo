package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokemonName ...string) error {
	if len(pokemonName) != 1 {
		return errors.New("you need to enter a pokemon name in you pokedex")
	}
	val, ok := cfg.pokemonCatched[pokemonName[0]]
	if !ok {
		return errors.New("You haven't caught this pokemon yet")
	}
	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Height: %v\n", val.Height)
	fmt.Printf("Weight: %v\n", val.Weight)
	fmt.Println("Stats:")
	for _, stat := range val.Stats {
		fmt.Printf("-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range val.Types {
		fmt.Printf("- %v\n", typ.Type.Name)
	}

	return nil
}
