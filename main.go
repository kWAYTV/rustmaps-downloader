package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

// Add constants for better maintainability
const (
	baseAPIURL     = "https://api.rustmaps.com/v4/maps/filter"
	outputDir      = "maps"
	requestTimeout = 10 * time.Second
	rateLimit      = time.Second
)

// Response represents the API response structure
type Response struct {
	Meta Meta  `json:"meta"`
	Data []Map `json:"data"`
}

// Meta contains pagination and status information
type Meta struct {
	Status       string   `json:"status"`
	StatusCode   int      `json:"statusCode"`
	Errors       []string `json:"errors,omitempty"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
	TotalItems   int      `json:"totalItems"`
	LastPage     bool     `json:"lastPage"`
}

// Map represents a single map entry
type Map struct {
	MapID string `json:"mapId"`
	Seed  int    `json:"seed"`
	Size  int    `json:"size"`
	URL   string `json:"url"`
}

func loadConfig() (string, string, error) {
	if err := godotenv.Load(); err != nil {
		return "", "", fmt.Errorf("error loading .env file: %w", err)
	}

	apiKey := os.Getenv("RUSTMAPS_API_KEY")
	filterID := os.Getenv("RUSTMAPS_FILTER_ID")

	if apiKey == "" || filterID == "" {
		return "", "", fmt.Errorf("RUSTMAPS_API_KEY and RUSTMAPS_FILTER_ID must be set in .env file")
	}

	// Create maps directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", "", fmt.Errorf("error creating maps directory: %w", err)
	}

	return apiKey, filterID, nil
}

func fetchMaps(apiKey, filterID string) ([]Map, error) {
	url := fmt.Sprintf("%s/%s", baseAPIURL, filterID)
	client := &http.Client{Timeout: requestTimeout}
	var allMaps []Map
	page := 0

	// Pre-allocate slice with reasonable capacity
	allMaps = make([]Map, 0, 1000)

	for {
		select {
		case <-time.After(rateLimit):
			response, err := fetchPage(client, url, apiKey, page)
			if err != nil {
				return nil, err
			}

			allMaps = append(allMaps, response.Data...)
			fmt.Printf("Fetched page %d, got %d maps. Total so far: %d\n",
				page, len(response.Data), len(allMaps))

			if response.Meta.LastPage {
				return allMaps, nil
			}
			page++
		}
	}
}

// Split page fetching logic into separate function
func fetchPage(client *http.Client, baseURL, apiKey string, page int) (*Response, error) {
	url := fmt.Sprintf("%s?page=%d&staging=false", baseURL, page)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &response, nil
}

func saveMapsToFile(maps []Map, filterID string) error {
	filename := filepath.Join(outputDir, fmt.Sprintf("rust_maps_%s.json", filterID))

	// Create backup of existing file if it exists
	if _, err := os.Stat(filename); err == nil {
		backupName := filepath.Join(outputDir, fmt.Sprintf("rust_maps_%s_%s.backup.json",
			filterID, time.Now().Format("20060102_150405")))
		if err := os.Rename(filename, backupName); err != nil {
			return fmt.Errorf("error creating backup: %w", err)
		}
	}

	data, err := json.MarshalIndent(maps, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	return os.WriteFile(filename, data, 0644)
}

func main() {
	apiKey, filterID, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	maps, err := fetchMaps(apiKey, filterID)
	if err != nil {
		log.Fatalf("Failed to fetch maps: %v", err)
	}

	if err := saveMapsToFile(maps, filterID); err != nil {
		log.Fatalf("Failed to save maps: %v", err)
	}

	fmt.Printf("\nTotal maps collected: %d\n", len(maps))
	fmt.Printf("Data saved to %s\n", fmt.Sprintf("rust_maps_%s.json", filterID))
}
