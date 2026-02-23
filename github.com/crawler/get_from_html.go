package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHtml(html string) string {
	body := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return ""
	}
	h1 := doc.Find("H1").First().Text()

	return strings.TrimSpace(h1)
}

func getFirstParagraphFromHTML(html string) string {
	body := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return ""
	}
	main := doc.Find("main")
	var p string
	if main.Length() != 0 {
		p = doc.Find("main").Find("p").First().Text()
	} else {
		p = doc.Find("p").First().Text()
	}

	return strings.TrimSpace(p)
}
