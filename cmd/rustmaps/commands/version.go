package commands

import (
	"github.com/spf13/cobra"
)

var Version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number of RustMaps CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("RustMaps CLI v%s", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
