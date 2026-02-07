package school_news

import (
	"fmt"
)

// ForceRefresh forces a new fetch of the news data.
func ForceRefresh() error {
	return globalStore.forceRefresh()
}

// GetLastRefreshTime returns the last refresh time in "YYYY/MM/DD HH:MM:SS" format.
func GetLastRefreshTime() string {
	return globalStore.getLastRefreshTime()
}

// CountNews returns the total number of news items.
func CountNews() (int, error) {
	items, err := globalStore.getItems()
	if err != nil {
		return 0, fmt.Errorf("failed to get items: %w", err)
	}
	return len(items), nil
}

// CountNewsByCategory returns the number of news items in a specific category.
func CountNewsByCategory(category string) (int, error) {
	items, err := globalStore.getItems()
	if err != nil {
		return 0, fmt.Errorf("failed to get items: %w", err)
	}

	count := 0
	for _, item := range items {
		if item.Category == category {
			count++
		}
	}
	return count, nil
}

// CountNewsByPublisher returns the number of news items by a specific publisher.
func CountNewsByPublisher(publisher string) (int, error) {
	items, err := globalStore.getItems()
	if err != nil {
		return 0, fmt.Errorf("failed to get items: %w", err)
	}

	count := 0
	for _, item := range items {
		if item.Publisher == publisher {
			count++
		}
	}
	return count, nil
}

// GetNews returns news items within the specified index range (start inclusive, end exclusive).
func GetNews(start, end int) ([]NewsItem, error) {
	items, err := globalStore.getItems()
	if err != nil {
		return nil, fmt.Errorf("failed to get items: %w", err)
	}

	return sliceItems(items, start, end), nil
}

// GetNewsByCategory returns news items in a specific category within the specified index range (start inclusive, end exclusive).
func GetNewsByCategory(category string, start, end int) ([]NewsItem, error) {
	items, err := globalStore.getItems()
	if err != nil {
		return nil, fmt.Errorf("failed to get items: %w", err)
	}

	var filtered []NewsItem
	for _, item := range items {
		if item.Category == category {
			filtered = append(filtered, item)
		}
	}

	return sliceItems(filtered, start, end), nil
}

// GetNewsByPublisher returns news items by a specific publisher within the specified index range (start inclusive, end exclusive).
func GetNewsByPublisher(publisher string, start, end int) ([]NewsItem, error) {
	items, err := globalStore.getItems()
	if err != nil {
		return nil, fmt.Errorf("failed to get items: %w", err)
	}

	var filtered []NewsItem
	for _, item := range items {
		if item.Publisher == publisher {
			filtered = append(filtered, item)
		}
	}

	return sliceItems(filtered, start, end), nil
}

// sliceItems helper function to handle pagination safely.
func sliceItems(items []NewsItem, start, end int) []NewsItem {
	total := len(items)
	if start < 0 {
		start = 0
	}
	if start >= total {
		return []NewsItem{}
	}
	if end > total {
		end = total
	}
	if start >= end {
		return []NewsItem{}
	}
	return items[start:end]
}
