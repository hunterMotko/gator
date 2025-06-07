package cmd

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/huntermotko/gator/internal/database"
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

func scrapeFeeds(s *State) error {
	// Get the next feed to fetch from the DB.
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}
	// Mark it as fetched.
	if err := s.db.MarkFeedFetched(ctx, feed.ID); err != nil {
		return err
	}
	// Fetch the feed using the URL (we already wrote this function)
	// Iterate over the items in the feed and print their titles to the console.
	rss, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}
	for _, item := range rss.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		_, err := s.db.CreatePost(ctx, database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    feed.ID,
			Title:     item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url:         item.Link,
			PublishedAt: publishedAt,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			fmt.Printf("Couldn't create post: %v\n", err)
			continue
		}
	}
	fmt.Printf("Feed %s collected, %v posts found\n", feed.Name, len(rss.Channel.Item))
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var feed RSSFeed
	if err := xml.Unmarshal(b, &feed); err != nil {
		return nil, err
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}
	return &feed, nil
}

func printFeed(feed *RSSFeed) {
	fmt.Println(feed.Channel.Title)
	fmt.Println(feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
		fmt.Println(item.Description)
		fmt.Println(item.Link)
		fmt.Println(item.PubDate)
	}
}
