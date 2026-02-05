package main

import "fmt"

func commandExplore(cfg *config, area ...string) error {

	if len(area) < 1 {
		fmt.Println("Please choose area you want to explore")
	}

	pokelist, err := cfg.pokeapiClient.ListPokeArea(area[0])
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
