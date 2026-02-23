package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	body := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return []string{}, err
	}
	var res []string

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		attr, ok := s.Attr("href")
		if !ok {
			return
		}
		attr = strings.TrimSpace(attr)
		if attr == "" {
			return
		}
		parsedRelative, err := url.Parse(attr)
		if err != nil {
			return
		}
		res = append(res, baseURL.ResolveReference(parsedRelative).String())

	})
	return res, nil
}
