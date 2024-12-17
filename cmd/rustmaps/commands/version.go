package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number of RustMaps CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("RustMaps CLI v%s\n", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
