package school_news

import "time"

// NewsItem represents a single news entry.
type NewsItem struct {
	Title       string
	Description string
	Link        string
	PubDate     time.Time
	Category    string
	Publisher   string
}
