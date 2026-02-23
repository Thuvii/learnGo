package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	html := strings.NewReader(htmlBody)

	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return []string{}, nil
	}

	var res []string

	doc.Find("img[src]").Each(func(i int, s *goquery.Selection) {
		attr, ok := s.Attr("src")
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
