package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type MapData struct {
	MapID string `json:"mapId"`
	Seed  int    `json:"seed"`
	Size  int    `json:"size"`
	URL   string `json:"url"`
}

type WorldSeed struct {
	Seed int `yaml:"seed"`
	Size int `yaml:"size"`
}

var updateConfigCmd = &cobra.Command{
	Use:   "update-config <maps.json> <config.yml>",
	Short: "Update config.yml with seeds from maps.json",
	Long: `Update a Rust Wipe Bot config file with seeds from a downloaded maps JSON file.
	
Example usage:
  # Using relative paths
  rustmaps update-config maps/rust_maps_filter123.json config.yml

  # Using absolute paths (wrap in quotes if paths contain spaces)
  rustmaps update-config "C:\Path\To\maps.json" "C:\Path\To\config.yml"`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		mapsFile := args[0]
		configFile := args[1]

		log.Info("📖 Reading maps file...")
		// Read maps JSON
		mapsData, err := ioutil.ReadFile(mapsFile)
		if err != nil {
			return fmt.Errorf("failed to read maps file: %w", err)
		}

		var maps []MapData
		if err := json.Unmarshal(mapsData, &maps); err != nil {
			return fmt.Errorf("failed to parse maps JSON: %w", err)
		}
		log.Infof("✅ Found %d maps in JSON file", len(maps))

		log.Info("📖 Reading config file...")
		// Read config YAML
		configData, err := ioutil.ReadFile(configFile)
		if err != nil {
			return fmt.Errorf("failed to read config file: %w", err)
		}

		// Split the file into lines to preserve comments
		lines := strings.Split(string(configData), "\n")
		var configLines []string
		inWorldSeeds := false

		// Process each line
		for _, line := range lines {
			if strings.Contains(line, "world_seeds:") {
				inWorldSeeds = true
				configLines = append(configLines, line)
				// Add new world seeds
				for _, m := range maps {
					configLines = append(configLines, fmt.Sprintf("  - seed: %d", m.Seed))
					configLines = append(configLines, fmt.Sprintf("    size: %d", m.Size))
				}
			} else if inWorldSeeds {
				if strings.HasPrefix(strings.TrimSpace(line), "-") || strings.HasPrefix(strings.TrimSpace(line), "seed:") || strings.HasPrefix(strings.TrimSpace(line), "size:") {
					continue
				}
				inWorldSeeds = false
				configLines = append(configLines, line)
			} else {
				configLines = append(configLines, line)
			}
		}

		log.Info("💾 Writing updated config...")
		// Write the updated config
		output := strings.Join(configLines, "\n")
		if err := ioutil.WriteFile(configFile, []byte(output), 0644); err != nil {
			return fmt.Errorf("failed to write updated config: %w", err)
		}

		log.Infof("✨ Successfully updated %s with %d seeds", configFile, len(maps))
		return nil
	},
}

func init() {
	// Configure logger
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
	RootCmd.AddCommand(updateConfigCmd)
}
