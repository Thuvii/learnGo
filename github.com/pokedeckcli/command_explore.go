package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, area ...string) error {

	if len(area) < 1 {
		return errors.New("Please choose area you want to explore")
	}

	pokelist, err := cfg.pokeapiClient.ListPokebyArea(area[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", area[0])
	fmt.Println("Found Pokemon: ")
	for _, poke := range pokelist.PokemonEncounters {
		fmt.Println(poke.Pokemon.Name)
	}
	return nil
}
