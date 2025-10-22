package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/bc0d3/trident-recon/pkg/config"
	"github.com/bc0d3/trident-recon/pkg/executor"
	"github.com/bc0d3/trident-recon/pkg/generator"
	"github.com/bc0d3/trident-recon/pkg/utils"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Generate and execute commands in tmux sessions",
	Long: `Generate reconnaissance commands and execute them in background tmux sessions.

This will create tmux sessions for each command and save session metadata.
You can monitor sessions using 'tmux attach' or 'trident-recon list'.

Examples:
  trident-recon run -u http://example.com
  trident-recon run -u http://example.com -o ~/scans/target1
  trident-recon run -l targets.txt --tools ffuf,gobuster`,
	RunE: runRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) error {
	// Validate flags
	if err := validateTargetFlags(); err != nil {
		return err
	}

	// Load config
	utils.PrintInfo("Loading configuration...")
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'trident-recon init' first)", err)
	}

	// Validate config
	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	// Get state directory
	stateDir := config.GetStateDir()

	// Get targets
	targets, err := getTargets()
	if err != nil {
		return err
	}

	utils.PrintSuccess(fmt.Sprintf("Found %d target(s)", len(targets)))
	fmt.Println()

	// Process each target
	for i, target := range targets {
		utils.PrintInfo(fmt.Sprintf("[%d/%d] Processing target: %s", i+1, len(targets), target))

		if err := runForTarget(cfg, target, stateDir); err != nil {
			utils.PrintError(fmt.Sprintf("Failed to run for %s: %v", target, err))
			continue
		}

		fmt.Println()
	}

	return nil
}

func runForTarget(cfg *config.Config, target, stateDir string) error {
	// Determine output directory
	var outDir string
	if outputDir != "" {
		outDir = utils.ExpandPath(outputDir)
	} else {
		// Parse URL to get domain
		_, domain, err := utils.ParseURL(target)
		if err != nil {
			return fmt.Errorf("invalid URL: %w", err)
		}
		domain = utils.SanitizeDomain(domain)

		// Generate timestamped directory
		baseDir := utils.ExpandPath(cfg.Global.OutputDir)
		outDir = utils.GenerateOutputDir(baseDir, domain)
	}

	// Create output directory
	if err := utils.EnsureDir(outDir); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	utils.PrintSuccess(fmt.Sprintf("Output directory: %s", outDir))

	// Generate commands
	utils.PrintInfo("Generating commands...")
	gen := generator.New(cfg, target, outDir)
	sessions, err := gen.Generate(toolsFilter, skipTools)
	if err != nil {
		return fmt.Errorf("failed to generate commands: %w", err)
	}

	utils.PrintSuccess(fmt.Sprintf("Generated %d command(s)", len(sessions)))

	// Save markdown file
	mdGen := generator.MarkdownGenerator{
		Target:    target,
		OutputDir: outDir,
		Sessions:  sessions,
	}

	markdown := mdGen.Generate()
	mdPath := filepath.Join(outDir, "commands.md")
	if err := utils.WriteFile(mdPath, markdown); err != nil {
		return fmt.Errorf("failed to write markdown: %w", err)
	}

	utils.PrintSuccess(fmt.Sprintf("Markdown saved to: %s", mdPath))

	// Generate plain text commands
	txtGen := generator.PlainTextGenerator{
		Sessions: sessions,
	}

	plainText := txtGen.Generate()
	txtPath := filepath.Join(outDir, "commands.txt")
	if err := utils.WriteFile(txtPath, plainText); err != nil {
		return fmt.Errorf("failed to write plain text commands: %w", err)
	}

	utils.PrintSuccess(fmt.Sprintf("Plain text commands saved to: %s", txtPath))
	fmt.Println()

	// Execute sessions
	utils.PrintInfo("Executing commands in tmux sessions...")
	exec := executor.NewExecutor(stateDir)

	// Validate sessions before execution
	if err := exec.ValidateSessions(sessions); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Execute all sessions
	successful, err := exec.ExecuteAll(sessions)
	if err != nil {
		return fmt.Errorf("execution failed: %w", err)
	}

	fmt.Println()
	utils.PrintSuccess(fmt.Sprintf("Successfully started %d/%d sessions", successful, len(sessions)))
	fmt.Println()
	fmt.Println("📋 Session Management:")
	fmt.Println("   List sessions:      trident-recon list")
	fmt.Println("   Attach to session:  tmux attach -t <session-name>")
	fmt.Println("   Kill all sessions:  trident-recon kill-all")
	fmt.Println()
	fmt.Printf("📁 Output directory:           %s\n", outDir)
	fmt.Printf("📄 Markdown file:              %s\n", mdPath)
	fmt.Printf("📄 Commands file (copy-paste): %s\n", txtPath)

	return nil
}
