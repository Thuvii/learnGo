package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type crawlerClient struct {
	httpClient *http.Client
}

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}

	client := &crawlerClient{httpClient: &http.Client{}}
	req.Header.Set("User-Agent", "BootCrawler/1.0")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error code %v", resp.StatusCode)
	}

	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("Wrong content-type: %v. Expected text/html", resp.Header.Get("content-type"))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil

}
