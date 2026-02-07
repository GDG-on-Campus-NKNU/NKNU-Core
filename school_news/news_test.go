package school_news

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// Sample RSS XML for testing
const sampleRSS = `<?xml version="1.0" encoding="utf-8"?>
<rss version="2.0">
<channel>
	<item>
		<title>【教務處公告】News 1</title>
		<description>Desc 1</description>
		<link>http://example.com/1</link>
		<pubDate>2026-02-07 10:00:00</pubDate>
		<category>Academic</category>
	</item>
	<item>
		<title>【總務處】News 2</title>
		<description>Desc 2</description>
		<link>http://example.com/2</link>
		<pubDate>Invalid Date</pubDate>
		<category>General</category>
	</item>
	<item>
		<title>News 3 No Publisher</title>
		<category>Other</category>
	</item>
</channel>
</rss>`

const malformedXML = `<rss><channel><item`

func TestFetchRSS(t *testing.T) {
	// Setup mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprint(w, sampleRSS)
	}))
	defer server.Close()

	// Override URL
	originalURL := rssURL
	rssURL = server.URL
	defer func() { rssURL = originalURL }()

	// Test fetch
	items, err := fetchRSS()
	if err != nil {
		t.Fatalf("fetchRSS failed: %v", err)
	}

	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}

	// Verify Item 1 (Normal)
	if items[0].Title != "【教務處公告】News 1" {
		t.Errorf("Item 1 Title mismatch")
	}
	if items[0].Publisher != "教務處" { // "公告" suffix removed
		t.Errorf("Item 1 Publisher extraction failed, got '%s'", items[0].Publisher)
	}
	expectedDate, _ := time.Parse("2006-01-02 15:04:05", "2026-02-07 10:00:00")
	if !items[0].PubDate.Equal(expectedDate) {
		t.Errorf("Item 1 Date mismatch")
	}

	// Verify Item 2 (Invalid Date, Publisher no suffix)
	if items[1].Publisher != "總務處" {
		t.Errorf("Item 2 Publisher mismatch, got '%s'", items[1].Publisher)
	}
	if !items[1].PubDate.IsZero() {
		t.Errorf("Item 2 Date should be zero for invalid date string")
	}

	// Verify Item 3 (No Publisher)
	if items[2].Publisher != "" {
		t.Errorf("Item 3 Publisher should be empty")
	}
}

func TestFetchRSSError(t *testing.T) {
	// 1. Network Error
	// Point to closed port
	originalURL := rssURL
	rssURL = "http://localhost:0" // invalid port usually
	defer func() { rssURL = originalURL }()

	_, err := fetchRSS()
	if err == nil {
		t.Error("fetchRSS should fail on network error")
	}

	// 2. HTTP Error Status
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	rssURL = server.URL

	_, err = fetchRSS()
	if err == nil || !strings.Contains(err.Error(), "bad status code") {
		t.Error("fetchRSS should fail on bad status code")
	}

	// 3. Malformed XML
	serverXML := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, malformedXML)
	}))
	defer serverXML.Close()
	rssURL = serverXML.URL

	_, err = fetchRSS()
	if err == nil || !strings.Contains(err.Error(), "failed to unmarshal XML") {
		t.Error("fetchRSS should fail on malformed XML")
	}
}

func TestStoreRefreshAndPublicAPI(t *testing.T) {
	// Setup mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, sampleRSS)
	}))
	defer server.Close()

	originalURL := rssURL
	rssURL = server.URL
	defer func() { rssURL = originalURL }()

	// Reset global store
	globalStore = &store{}

	// Test ForceRefresh
	err := ForceRefresh()
	if err != nil {
		t.Fatalf("ForceRefresh failed: %v", err)
	}

	// Test GetLastRefreshTime
	lastTime := GetLastRefreshTime()
	if lastTime == "" {
		t.Error("GetLastRefreshTime should not be empty")
	}

	// Test CountNews
	count, err := CountNews()
	if err != nil {
		t.Errorf("CountNews error: %v", err)
	}
	if count != 3 {
		t.Errorf("CountNews expected 3, got %d", count)
	}

	// Test CountNewsByCategory
	cCat, err := CountNewsByCategory("Academic")
	if err != nil || cCat != 1 {
		t.Errorf("CountNewsByCategory Academic expected 1, got %d", cCat)
	}
	cCat, _ = CountNewsByCategory("NonExistent")
	if cCat != 0 {
		t.Errorf("CountNewsByCategory NonExistent expected 0")
	}

	// Test CountNewsByPublisher
	cPub, err := CountNewsByPublisher("教務處")
	if err != nil || cPub != 1 {
		t.Errorf("CountNewsByPublisher expected 1, got %d", cPub)
	}

	// Test GetNews pagination
	// Case 1: All
	all, err := GetNews(0, 10)
	if err != nil || len(all) != 3 {
		t.Errorf("GetNews(0, 10) expected 3 items")
	}
	// Case 2: Subset
	subset, err := GetNews(1, 2)
	if err != nil || len(subset) != 1 {
		t.Errorf("GetNews(1, 2) expected 1 item")
	} else if subset[0].Title != "【總務處】News 2" {
		t.Errorf("GetNews(1, 2) content mismatch")
	}
	// Case 3: Out of bounds (Start >= Total)
	empty, err := GetNews(10, 20)
	if err != nil || len(empty) != 0 {
		t.Errorf("GetNews(10, 20) expected empty")
	}
	// Case 4: Invalid start
	invalidStart, err := GetNews(-1, 2)
	if err != nil || len(invalidStart) != 2 {
		t.Errorf("GetNews(-1, 2) expected start corrected to 0, got len %d", len(invalidStart))
	}
	// Case 5: Start >= End
	invOrder, _ := GetNews(2, 2)
	if len(invOrder) != 0 {
		t.Errorf("GetNews(2, 2) expected empty")
	}

	// Test GetNewsByCategory
	catItems, err := GetNewsByCategory("Academic", 0, 5)
	if err != nil || len(catItems) != 1 {
		t.Errorf("GetNewsByCategory expected 1 item")
	}
	// Test pagination on filtered
	catItemsEmpty, _ := GetNewsByCategory("Academic", 1, 5)
	if len(catItemsEmpty) != 0 {
		t.Errorf("GetNewsByCategory offset 1 expected empty result for 1 item list")
	}

	// Test GetNewsByPublisher
	pubItems, err := GetNewsByPublisher("教務處", 0, 5)
	if err != nil || len(pubItems) != 1 {
		t.Errorf("GetNewsByPublisher expected 1 item")
	}
}

func TestStoreEnsureData(t *testing.T) {
	// Setup mock server
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		fmt.Fprint(w, sampleRSS)
	}))
	defer server.Close()

	originalURL := rssURL
	rssURL = server.URL
	defer func() { rssURL = originalURL }()

	// Reset global store with expired time
	globalStore = &store{
		items:       []NewsItem{}, // empty
		lastUpdated: time.Time{},  // zero
	}

	// 1. First call should trigger fetch
	globalStore.ensureData()
	if requestCount != 1 {
		t.Errorf("Expected 1 request, got %d", requestCount)
	}

	// 2. Second call should cache hit (interval not passed)
	globalStore.ensureData()
	if requestCount != 1 {
		t.Errorf("Expected optimization (no new request), got %d requests", requestCount)
	}

	// 3. Force expire (simulate time passing)
	globalStore.mu.Lock()
	globalStore.lastUpdated = time.Now().Add(-10 * time.Minute)
	globalStore.mu.Unlock()

	globalStore.ensureData()
	if requestCount != 2 {
		t.Errorf("Expected 2 requests after expiration, got %d", requestCount)
	}
}

func TestStoreErrorPropagation(t *testing.T) {
	// Simulate failing server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	originalURL := rssURL
	rssURL = server.URL
	defer func() { rssURL = originalURL }()

	globalStore = &store{} // clear store

	// Verify all exported functions propagate error
	if _, err := CountNews(); err == nil {
		t.Error("CountNews expected error")
	}
	if _, err := CountNewsByCategory("A"); err == nil {
		t.Error("CountNewsByCategory expected error")
	}
	if _, err := CountNewsByPublisher("A"); err == nil {
		t.Error("CountNewsByPublisher expected error")
	}
	if _, err := GetNews(0, 10); err == nil {
		t.Error("GetNews expected error")
	}
	if _, err := GetNewsByCategory("A", 0, 10); err == nil {
		t.Error("GetNewsByCategory expected error")
	}
	if _, err := GetNewsByPublisher("A", 0, 10); err == nil {
		t.Error("GetNewsByPublisher expected error")
	}
}

func TestSliceLogicSafety(t *testing.T) {
	// Direct test of slice logic with edge cases just to be sure
	items := []NewsItem{{Title: "A"}, {Title: "B"}, {Title: "C"}}

	// Start > End (checked explicitly in GetNews case 5 too, but here direct)
	res := sliceItems(items, 5, 2)
	if len(res) != 0 {
		t.Error("Start > End should return empty")
	}

	// Start == End (Empty range)
	res = sliceItems(items, 1, 1)
	if len(res) != 0 {
		t.Error("Start == End should return empty")
	}

	// End > Total -> capped
	res = sliceItems(items, 0, 100)
	if len(res) != 3 {
		t.Error("End > Total should return all")
	}

	// Start < 0 -> 0
	res = sliceItems(items, -5, 1)
	if len(res) != 1 {
		t.Error("Negative start should become 0")
	}
}
