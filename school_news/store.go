package school_news

import (
	"fmt"
	"sync"
	"time"
)

const (
	refreshInterval = 5 * time.Minute
)

// store manages the cached news items.
type store struct {
	items       []NewsItem
	lastUpdated time.Time
	mu          sync.RWMutex
}

var globalStore = &store{}

// ensureData checks if data needs to be sorted/fetched.
func (s *store) ensureData() error {
	s.mu.RLock()
	isValid := !s.lastUpdated.IsZero() && time.Since(s.lastUpdated) < refreshInterval && len(s.items) > 0
	s.mu.RUnlock()

	if isValid {
		return nil
	}

	// Double-checked locking
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.lastUpdated.IsZero() && time.Since(s.lastUpdated) < refreshInterval && len(s.items) > 0 {
		return nil
	}

	return s.refresh()
}

// refresh performs the actual update. Caller must hold the write lock.
func (s *store) refresh() error {
	items, err := fetchRSS()
	if err != nil {
		return fmt.Errorf("failed to fetch news: %w", err)
	}

	s.items = items
	s.lastUpdated = time.Now()
	return nil
}

// forceRefresh forces a new fetch of the news data.
func (s *store) forceRefresh() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.refresh()
}

// getLastRefreshTime returns the formatted last refresh time.
func (s *store) getLastRefreshTime() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.lastUpdated.IsZero() {
		return ""
	}
	return s.lastUpdated.Format("2006/01/02 15:04:05")
}

// getItems returns a thread-safe copy of the items.
func (s *store) getItems() ([]NewsItem, error) {
	if err := s.ensureData(); err != nil {
		return nil, err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a copy to be safe
	result := make([]NewsItem, len(s.items))
	copy(result, s.items)
	return result, nil
}
