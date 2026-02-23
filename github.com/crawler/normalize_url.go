package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(fullUrl string) (string, error) {
	parsedURL, err := url.Parse(fullUrl)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := parsedURL.Host + parsedURL.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil

}
