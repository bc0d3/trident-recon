package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	targetURL    string
	targetList   string
	outputDir    string
	generateOnly bool
	runCommands  bool
	toolsFilter  []string
	skipTools    []string
	toolFilter   string
	version      string
	commit       string
	date         string
)

var rootCmd = &cobra.Command{
	Use:   "trident-recon",
	Short: "🔱 Bug bounty tool orchestrator",
	Long: `Trident Recon - Malleable reconnaissance tool orchestrator

Generate and execute multiple recon tools in tmux background sessions.
Manage your bug bounty workflow with ease.`,
	Version: "1.0.0",
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

// SetVersionInfo sets version information
func SetVersionInfo(v, c, d string) {
	version = v
	commit = c
	date = d
	rootCmd.Version = fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date)
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&targetURL, "url", "u", "", "Target URL")
	rootCmd.PersistentFlags().StringVarP(&targetList, "list", "l", "", "File with list of targets")
	rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "", "Output directory")
	rootCmd.PersistentFlags().BoolVarP(&generateOnly, "generate", "g", false, "Generate commands only (don't execute)")
	rootCmd.PersistentFlags().BoolVarP(&runCommands, "run", "r", false, "Generate and run commands")
	rootCmd.PersistentFlags().StringSliceVarP(&toolsFilter, "tools", "t", nil, "Run only specified tools (comma-separated)")
	rootCmd.PersistentFlags().StringSliceVar(&skipTools, "skip", nil, "Skip specified tools (comma-separated)")
}

func validateTargetFlags() error {
	if targetURL == "" && targetList == "" {
		return fmt.Errorf("either --url or --list must be specified")
	}
	if targetURL != "" && targetList != "" {
		return fmt.Errorf("cannot specify both --url and --list")
	}
	return nil
}

func validateActionFlags() error {
	if !generateOnly && !runCommands {
		return fmt.Errorf("either --generate or --run must be specified")
	}
	if generateOnly && runCommands {
		return fmt.Errorf("cannot specify both --generate and --run")
	}
	return nil
}
