package main

import "net/url"

type PageData struct {
	URL            string   `json:"url"`
	H1             string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
}

func extractPageData(html, pageURL string) PageData {
	url, err := url.Parse(pageURL)
	H1 := getH1FromHtml(html)

	firstP := getFirstParagraphFromHTML(html)
	if err != nil {
		return PageData{
			URL:            pageURL,
			H1:             H1,
			FirstParagraph: firstP,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
		}
	}

	link, err := getURLsFromHTML(html, url)
	if err != nil {
		return PageData{}
	}

	imageUrls, err := getImagesFromHTML(html, url)
	if err != nil {
		return PageData{}
	}

	pageData := &PageData{}

	pageData.URL = pageURL
	pageData.H1 = H1
	pageData.FirstParagraph = firstP
	pageData.OutgoingLinks = link
	pageData.ImageURLs = imageUrls

	return *pageData

}
