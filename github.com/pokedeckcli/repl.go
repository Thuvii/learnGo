package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
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

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		userInput := scanner.Text()
		userInputClean := cleanInput(userInput)
	  commandName := userInputClean[0]



		command,exist := getCommand()[commandName]

		if exist{
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
		continue
		}else{
		fmt.Println("Unknow command")
		}
	}
}


func getCommand() map[string]cliCommand{
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
  	}
}

