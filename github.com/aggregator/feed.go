package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"html"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Thuvii/aggregator/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{}

	res, err := http.NewRequestWithContext(context.Background(), "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	res.Header.Set("User-Agent", "gator")

	resp, err := client.Do(res)
	if err != nil {
		return &RSSFeed{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	var rssFeed RSSFeed
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return &RSSFeed{}, err
	}

	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)

	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
	}

	return &rssFeed, nil

}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Print(err)
		return
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Print(err)
		return
	}

	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Print(err)
		return
	}
	for _, item := range rss.Channel.Item {
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		valid := true
		if err != nil {
			t, err = time.Parse(time.RFC1123, item.PubDate)
			if err != nil {
				valid = false
				log.Printf("could not parse published_at: %v", err)
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: sql.NullTime{Time: t, Valid: valid},
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("could not create post: %v", err)
		}

	}

}
