package main

import (
	"fmt"
	"os"
	"time"
)

func commandExit(ac *config, area ...string) error {
	loadingString := "...\n"
	for _, r := range loadingString {
		fmt.Print(string(r))
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Print("Closing the Pokedex...Goodbye!")
	os.Exit(0)
	return nil
}
