package main

import (
	"fmt"
	"os"

	"github.com/bc0d3/trident-recon/cmd"
)

var (
	// Version info (set by ldflags during build)
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func main() {
	// Set version info
	cmd.SetVersionInfo(Version, Commit, Date)

	// Execute command
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
