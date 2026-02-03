package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.ListLocation(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationResp.Next
	cfg.previousLocationUrl = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil

}

func commandMapb(cfg *config) error {
	if cfg.previousLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocation(cfg.previousLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationResp.Next
	cfg.previousLocationUrl = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil

}
