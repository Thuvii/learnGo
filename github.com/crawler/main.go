package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided! Example: go run . \"website to crawl\" \"max routine\" \"max pages\" ")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Print(err)
		return
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	config, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Print(err)
		return
	}
	config.wg.Add(1)
	go config.crawlPage(rawBaseURL)
	config.wg.Wait()

	for url, data := range config.pages {
		fmt.Printf("%s - %s\n", url, data.URL)
	}
	writeJSONReport(config.pages, "report.json")
}
