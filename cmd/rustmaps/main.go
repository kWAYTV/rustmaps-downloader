package main

import (
	"fmt"
	"os"

	"rustmaps-downloader/cmd/rustmaps/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
