package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, pokemonName ...string) error {
	if len(pokemonName) != 1 {
		return errors.New("Please pick a pokemon you want to catch\n")
	}
	pokeInfo, err := cfg.pokeapiClient.PokemonInfo(pokemonName[0])
	if err != nil {
		return err
	}

	catchRate := rand.IntN(100)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeInfo.Name)

	exp := pokeInfo.BaseExperience
	catchStatus := false
	if exp <= 100 {
		if catchRate > 15 {
			fmt.Printf("%s was caught!\n", pokeInfo.Name)
			catchStatus = true
		} else {
			fmt.Printf("%s escaped!\n", pokeInfo.Name)
		}
	} else if exp <= 200 && exp > 100 {
		if catchRate > 50 {
			fmt.Printf("%s was caught!\n", pokeInfo.Name)
			catchStatus = true
		} else {
			fmt.Printf("%s escaped!\n", pokeInfo.Name)

		}

	} else {
		if catchRate > 90 {
			fmt.Printf("%s was caught!\n", pokeInfo.Name)
			catchStatus = true
		} else {
			fmt.Printf("%s escaped!\n", pokeInfo.Name)
		}
	}

	if catchStatus {
		cfg.pokemonCatched[pokeInfo.Name] = pokeInfo
	}

	return nil
}
