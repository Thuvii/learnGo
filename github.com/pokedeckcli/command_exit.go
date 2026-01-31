package main

import(
  "fmt"
  "os"
  "time"
)

func commandExit() error {
	loadingString := "Closing the Pokedex... "
  for _, r := range loadingString {
		fmt.Print(string(r))
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Print("Goodbye!")
  os.Exit(0)
	return nil
}
