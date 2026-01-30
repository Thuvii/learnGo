package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	res := strings.Fields(lowerText)
	return res
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		userInput := scanner.Text()
		userInputClean := cleanInput(userInput)
		if userInputClean[0] == "exit" {

			commandMap["exit"].callback()
		}
	}
}

func commandExit() error {
	for i := 0; i < 5; i++ {
		dot := "."
		dot = dot + "."
		dotPrint := fmt.Sprintf(dot)
		fmt.Print(dotPrint)
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
