package main

import (
	"fmt"
)

func commandHelp(ac *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommand() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
