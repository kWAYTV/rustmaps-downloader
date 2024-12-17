package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const (
	baseAPIURL     = "https://api.rustmaps.com/v4/maps/filter"
	outputDir      = "maps"
	requestTimeout = 10 * time.Second
	rateLimit      = time.Second
)

type Response struct {
	Meta Meta  `json:"meta"`
	Data []Map `json:"data"`
}

type Meta struct {
	Status       string   `json:"status"`
	StatusCode   int      `json:"statusCode"`
	Errors       []string `json:"errors,omitempty"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
	TotalItems   int      `json:"totalItems"`
	LastPage     bool     `json:"lastPage"`
}

type Map struct {
	MapID string `json:"mapId"`
	Seed  int    `json:"seed"`
	Size  int    `json:"size"`
	URL   string `json:"url"`
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download maps based on filter",
	Long:  `Download maps from RustMaps API using the specified filter ID.`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	RootCmd.AddCommand(downloadCmd)
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
	baseFilename := filepath.Join(outputDir, fmt.Sprintf("rust_maps_%s.json", filterID))
	filename := baseFilename
	counter := 1

	// Find first available filename
	for {
		_, err := os.Stat(filename)
		if os.IsNotExist(err) {
			break
		}
		filename = filepath.Join(outputDir, fmt.Sprintf("rust_maps_%s_%d.json", filterID, counter))
		counter++
	}

	data, err := json.MarshalIndent(maps, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	fmt.Printf("Data saved to %s\n", filename)
	return nil
}
