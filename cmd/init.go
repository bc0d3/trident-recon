package cmd

import (
	"fmt"

	"github.com/bc0d3/trident-recon/pkg/config"
	"github.com/bc0d3/trident-recon/pkg/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize trident-recon configuration",
	Long: `Initialize trident-recon by creating default configuration files.

This will create:
  - Config file at ~/.config/trident-recon/config.yaml
  - State directory at ~/.local/state/trident-recon`,
	RunE: runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) error {
	configPath := config.GetConfigPath()
	stateDir := config.GetStateDir()

	// Check if config already exists
	if utils.FileExists(configPath) {
		confirm, err := utils.PromptConfirm(fmt.Sprintf("Config already exists at %s. Overwrite?", configPath))
		if err != nil {
			return err
		}
		if !confirm {
			utils.PrintInfo("Initialization cancelled")
			return nil
		}
	}

	// Create config directory
	utils.PrintInfo("Creating configuration directory...")
	if err := utils.WriteFile(configPath, config.DefaultConfig); err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}

	// Create state directory
	utils.PrintInfo("Creating state directory...")
	if err := utils.EnsureDir(stateDir); err != nil {
		return fmt.Errorf("failed to create state directory: %w", err)
	}

	utils.PrintSuccess("Configuration initialized successfully!")
	fmt.Println()
	fmt.Printf("📁 Config file: %s\n", configPath)
	fmt.Printf("📁 State directory: %s\n", stateDir)
	fmt.Println()
	utils.PrintInfo("Edit the config file to customize your wordlists and tools")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Edit config: vim " + configPath)
	fmt.Println("  2. Generate commands: trident-recon -u http://example.com -g")
	fmt.Println("  3. Run reconnaissance: trident-recon -u http://example.com -r")

	return nil
}
