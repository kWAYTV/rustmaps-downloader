package commands

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "rustmaps",
	Short: "RustMaps CLI tool for downloading and managing Rust maps",
	Long: `A CLI tool to interact with the RustMaps API.
Currently supports downloading maps based on filters.`,
}
