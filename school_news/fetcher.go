package school_news

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	rssURL = "https://news.nknu.edu.tw/nknu_News/RSS.ashx"
)

var (
	publisherRegex = regexp.MustCompile(`^【(.*?)】`)
)

// internal XML structs
type rssRoot struct {
	Channel rssChannel `xml:"channel"`
}

type rssChannel struct {
	Items []rssRawItem `xml:"item"`
}

type rssRawItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Category    string `xml:"category"`
}

// fetchRSS performs the actual network request and parsing.
func fetchRSS() ([]NewsItem, error) {
	resp, err := http.Get(rssURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var root rssRoot
	if err := xml.Unmarshal(data, &root); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	var items []NewsItem
	for _, raw := range root.Channel.Items {
		// Parse Date: 2026-02-07 20:42:30
		pubDate, err := time.Parse("2006-01-02 15:04:05", raw.PubDate)
		if err != nil {
			pubDate = time.Time{}
		}

		// Parse Publisher from Title
		publisher := ""
		matches := publisherRegex.FindStringSubmatch(raw.Title)
		if len(matches) > 1 {
			publisher = matches[1]
			// Heuristic: remove "公告" suffix if it's not the only word
			if len([]rune(publisher)) > 2 && strings.HasSuffix(publisher, "公告") {
				publisher = string([]rune(publisher)[:len([]rune(publisher))-2])
			}
		}

		items = append(items, NewsItem{
			Title:       raw.Title,
			Description: raw.Description,
			Link:        raw.Link,
			PubDate:     pubDate,
			Category:    raw.Category,
			Publisher:   publisher,
		})
	}

	return items, nil
}
