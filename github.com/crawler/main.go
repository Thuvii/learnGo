package main

import (
	"fmt"
	"net/url"
)

func main() {
	BASE_URL := "https://wagslane.dev"
	html, err := getHTML(BASE_URL)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	baseURL, err := url.Parse(BASE_URL)
	fmt.Println(getURLsFromHTML(html, baseURL))

}
