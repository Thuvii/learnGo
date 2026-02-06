package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Thuvii/pokedeckcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationUrl     *string
	previousLocationUrl *string
	pokemonCatched      map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	res := strings.Fields(lowerText)
	return res
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		userInput := scanner.Text()
		userInputClean := cleanInput(userInput)
		args := []string{}
		if len(userInputClean) == 0 {
			continue
		}
		if len(userInputClean) > 1 {
			args = userInputClean[1:]
		}
		commandName := userInputClean[0]

		command, exist := getCommand()[commandName]

		if exist {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknow command")
		}
	}
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display 20 next location",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display 20 previous location",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List pokemons from specific area",
			callback:    commandExplore,
		}, "catch": {
			name:        "catch",
			description: "catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect pokemon",
			callback:    commandInspect,
		}, "pokedex": {
			name:        "pokedex",
			description: "Display every pokemon have caught",
			callback:    commandPokedex,
		},
	}
}
