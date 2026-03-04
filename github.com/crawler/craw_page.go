package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}
	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}
	fmt.Printf("crawling %s\n", rawCurrentURL)
	data := extractPageData(htmlBody, rawCurrentURL)
	cfg.setPageData(normalizedURL, data)

	for _, nextURL := range data.OutgoingLinks {
		// fmt.Println("Crawling " + nextURL)
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
