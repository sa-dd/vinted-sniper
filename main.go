package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Vinted API response structures
type VintedResponse struct {
	Items []VintedItem `json:"items"`
}

type VintedItem struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Status    string `json:"status"`
	SizeTitle string `json:"size_title"`
	User      struct {
		Login string `json:"login"`
	} `json:"user"`
	Photo struct {
		URL string `json:"url"`
	} `json:"photo"`
	Price struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currency_code"`
	} `json:"price"`
}

// Configuration
var (
	discordToken   = "MTEzMjAyNTY1MzA5NzY2MDQ3OA.GSGaZJ.YB_dFAx6gul8SDngWZjeLTbyk8_eruttn4Zobg"
	discordChannel = "1377199803607420938"
	vintedURL      = "https://www.vinted.co.uk/api/v2/catalog/items?page=1&per_page=96&time=1748554179&search_text=&catalog_ids=&order=newest_first&catalog_from=0&size_ids=206,207,208,209,210,211,212,308,1624&brand_ids=53,14,535,8715,162,21099,345,245062,359177,484362,313669,378906,140618,1798422,597509,1065021,57144,345731,269830,99164,73458,719079,670432,299684,1985410,311812,291429,1037965,472855,511110,299838,8139,401801,3063,1412112,725037,164166,190014,46923,506331,13727,345562,335419,318349,276609,228018,45765,6526387,204980,4810534,296216,103684,205836,834670,559746,1195,24491,280255,672258,222762&status_ids=&color_ids=&material_ids="
	checkInterval  = 1 * time.Second
)

// State management
var (
	lastItemIDs  = make(map[int64]bool)
	stateMutex   sync.Mutex
	requestCount = 0 // Track number of requests made
)

func main() {
	if discordToken == "" || discordChannel == "" {
		log.Fatal("DISCORD_TOKEN and DISCORD_CHANNEL environment variables must be set")
	}

	// Create Discord session
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Start session
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}
	defer dg.Close()

	log.Println("Bot is now running. Press CTRL-C to exit.")

	// Start the Vinted polling
	go pollVintedAPI(dg)

	// Wait for termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func pollVintedAPI(s *discordgo.Session) {
	// Create HTTP client with keep-alive
	transport := &http.Transport{
		MaxIdleConns:        10,
		IdleConnTimeout:     90 * time.Second,
		DisableCompression:  false,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: false},
		MaxIdleConnsPerHost: 10,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	// Initial delay to let Discord connect
	time.Sleep(2 * time.Second)

	// Polling loop
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for range ticker.C {
		items, err := fetchVintedItems(client)
		if err != nil {
			log.Printf("Error fetching items: %v", err)
			continue
		}

		newItems := findNewItems(items)

		// Increment request counter
		requestCount++

		// Skip sending for first 2 requests
		if requestCount > 2 && len(newItems) > 0 {
			log.Printf("Found %d new items, sending to Discord", len(newItems))
			sendItemsToDiscord(s, newItems)
		} else if requestCount <= 2 {
			log.Printf("Skipping send for request #%d (initial setup)", requestCount)
		} else {
			log.Println("No new items found")
		}

		// Update state with current items
		updateItemState(items)
	}
}

func fetchVintedItems(client *http.Client) ([]VintedItem, error) {
	req, err := http.NewRequest("GET", vintedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	// Set required headers
	req.Header = http.Header{
		"authority":          {"www.vinted.co.uk"},
		"accept":             {"application/json, text/plain, */*"},
		"accept-language":    {"en-uk-fr"},
		"cache-control":      {"max-age=0"},
		"cookie":             {""}, // Get cookie from env
		"priority":           {"u=1, i"},
		"referer":            {"https://www.vinted.co.uk/catalog"},
		"sec-ch-ua":          {"\"Chromium\";v=\"136\", \"Microsoft Edge\";v=\"136\", \"Not.A/Brand\";v=\"99\""},
		"sec-ch-ua-mobile":   {"?0"},
		"sec-ch-ua-platform": {"\"Windows\""},
		"sec-fetch-dest":     {"empty"},
		"sec-fetch-mode":     {"cors"},
		"sec-fetch-site":     {"same-origin"},
		"user-agent":         {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Edg/136.0.0.0"},
		"x-anon-id":          {"a8c8ace7-ebf8-4144-9e8f-70a50132ec5a"},
		"x-csrf-token":       {"75f6c9fa-dc8e-4e52-a000-e09dd4084b3e"},
		"x-money-object":     {"true"},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response failed: %w", err)
	}

	var response VintedResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return response.Items, nil
}

func findNewItems(items []VintedItem) []VintedItem {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	var newItems []VintedItem
	for _, item := range items {
		if !lastItemIDs[item.ID] {
			newItems = append(newItems, item)
		} else {
			// Items are ordered newest first, so we can stop at first known item
			break
		}
	}
	return newItems
}

func updateItemState(items []VintedItem) {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	// Reset state
	lastItemIDs = make(map[int64]bool)

	// Add current items to state
	for _, item := range items {
		lastItemIDs[item.ID] = true
	}
}

func sendItemsToDiscord(s *discordgo.Session, items []VintedItem) {
	for _, item := range items {
		embed := createItemEmbed(item)
		_, err := s.ChannelMessageSendEmbed(discordChannel, embed)
		if err != nil {
			log.Printf("Failed to send message to Discord: %v", err)
		}
		// Avoid rate limits
		time.Sleep(500 * time.Millisecond)
	}
}

func createItemEmbed(item VintedItem) *discordgo.MessageEmbed {
	price := fmt.Sprintf("%s %s", item.Price.Amount, item.Price.CurrencyCode)

	return &discordgo.MessageEmbed{
		Title:       item.Title,
		URL:         item.URL,
		Description: "New item listed on Vinted!",
		Color:       0x1FC598, // Vinted green
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: item.Photo.URL,
		},
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Price", Value: price, Inline: true},
			{Name: "Condition", Value: item.Status, Inline: true},
			{Name: "Size", Value: item.SizeTitle, Inline: true},
			{Name: "Seller", Value: item.User.Login, Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Vinted Monitor",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
