package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/bc0d3/trident-recon/pkg/config"
	"github.com/bc0d3/trident-recon/pkg/generator"
	"github.com/bc0d3/trident-recon/pkg/utils"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate commands without executing",
	Long: `Generate reconnaissance commands and save to markdown file.

This will create a commands.md file with all the generated commands
that can be manually executed.

Examples:
  trident-recon generate -u http://example.com
  trident-recon generate -u http://example.com -o ~/scans/target1
  trident-recon generate -u http://example.com --tools ffuf,gobuster`,
	RunE: runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func runGenerate(cmd *cobra.Command, args []string) error {
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

	// Get targets
	targets, err := getTargets()
	if err != nil {
		return err
	}

	utils.PrintSuccess(fmt.Sprintf("Found %d target(s)", len(targets)))

	// If multiple targets, generate for all together (for domain list support)
	if len(targets) > 1 {
		return generateForMultipleTargets(cfg, targets)
	}

	// Single target - process individually
	for i, target := range targets {
		utils.PrintInfo(fmt.Sprintf("[%d/%d] Processing target: %s", i+1, len(targets), target))

		if err := generateForTarget(cfg, target, ""); err != nil {
			utils.PrintError(fmt.Sprintf("Failed to generate for %s: %v", target, err))
			continue
		}
	}

	return nil
}

func generateForTarget(cfg *config.Config, target string, domainListFile string) error {
	// Parse URL to get domain
	_, domain, err := utils.ParseURL(target)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	domain = utils.SanitizeDomain(domain)

	// Determine output directory
	var outDir string
	if outputDir != "" {
		// User specified output directory: outputDir/domain/
		outDir = filepath.Join(utils.ExpandPath(outputDir), domain)
	} else {
		// No output specified: use domain only
		baseDir := utils.ExpandPath(cfg.Global.OutputDir)
		outDir = filepath.Join(baseDir, domain)
	}

	// Create output directory
	if err := utils.EnsureDir(outDir); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate commands
	gen := generator.New(cfg, target, outDir)
	sessions, err := gen.Generate(toolsFilter, skipTools)
	if err != nil {
		return fmt.Errorf("failed to generate commands: %w", err)
	}

	utils.PrintSuccess(fmt.Sprintf("Generated %d command(s)", len(sessions)))

	// Generate markdown
	mdGen := generator.MarkdownGenerator{
		Target:    target,
		OutputDir: outDir,
		Sessions:  sessions,
	}

	markdown := mdGen.Generate()

	// Save markdown file
	mdPath := filepath.Join(outDir, "comandos.md")
	if err := utils.WriteFile(mdPath, markdown); err != nil {
		return fmt.Errorf("failed to write markdown: %w", err)
	}

	utils.PrintSuccess(fmt.Sprintf("Markdown saved to: %s", mdPath))

	// Generate plain text commands
	txtGen := generator.PlainTextGenerator{
		Sessions: sessions,
	}

	plainText := txtGen.Generate()

	// Save plain text file
	txtPath := filepath.Join(outDir, "comandos.txt")
	if err := utils.WriteFile(txtPath, plainText); err != nil {
		return fmt.Errorf("failed to write plain text commands: %w", err)
	}

	utils.PrintSuccess(fmt.Sprintf("Plain text commands saved to: %s", txtPath))
	fmt.Println()
	fmt.Println("📋 Quick Reference:")
	fmt.Printf("   Output directory: %s\n", outDir)
	fmt.Printf("   Markdown file: %s\n", mdPath)
	fmt.Printf("   Commands file (copy-paste): %s\n", txtPath)
	fmt.Printf("   Total commands: %d\n", len(sessions))
	fmt.Println()
	utils.PrintInfo("Review the files and execute manually or use 'trident-recon run'")

	return nil
}

func getTargets() ([]string, error) {
	if targetURL != "" {
		return []string{targetURL}, nil
	}

	if targetList != "" {
		targets, err := utils.ReadLines(targetList)
		if err != nil {
			return nil, fmt.Errorf("failed to read targets file: %w", err)
		}
		if len(targets) == 0 {
			return nil, fmt.Errorf("no targets found in file")
		}
		return targets, nil
	}

	return nil, fmt.Errorf("no target specified")
}

func generateForMultipleTargets(cfg *config.Config, targets []string) error {
	// Determine base output directory
	var baseOutDir string
	if outputDir != "" {
		baseOutDir = utils.ExpandPath(outputDir)
	} else {
		baseOutDir = utils.ExpandPath(cfg.Global.OutputDir)
	}

	// Create base output directory
	if err := utils.EnsureDir(baseOutDir); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create domains list file in base directory
	domainListFile := filepath.Join(baseOutDir, "domains.txt")
	var domains []string
	for _, target := range targets {
		_, domain, err := utils.ParseURL(target)
		if err != nil {
			utils.PrintError(fmt.Sprintf("Failed to parse %s: %v", target, err))
			continue
		}
		domains = append(domains, utils.SanitizeDomain(domain))
	}

	if err := utils.WriteLines(domainListFile, domains); err != nil {
		return fmt.Errorf("failed to write domains file: %w", err)
	}

	utils.PrintSuccess(fmt.Sprintf("Created domains file: %s", domainListFile))

	// Process each target in its own subdirectory
	for i, target := range targets {
		utils.PrintInfo(fmt.Sprintf("[%d/%d] Processing target: %s", i+1, len(targets), target))

		// Parse domain for this target
		_, domain, err := utils.ParseURL(target)
		if err != nil {
			utils.PrintError(fmt.Sprintf("Failed to parse %s: %v", target, err))
			continue
		}
		domain = utils.SanitizeDomain(domain)

		// Create subdirectory for this target
		targetOutDir := filepath.Join(baseOutDir, domain)
		if err := utils.EnsureDir(targetOutDir); err != nil {
			utils.PrintError(fmt.Sprintf("Failed to create directory for %s: %v", target, err))
			continue
		}

		// Generate commands for this target
		gen := generator.New(cfg, target, targetOutDir)
		gen.SetDomainListFile(domainListFile)

		sessions, err := gen.Generate(toolsFilter, skipTools)
		if err != nil {
			utils.PrintError(fmt.Sprintf("Failed to generate for %s: %v", target, err))
			continue
		}

		// Generate markdown for this target
		mdGen := generator.MarkdownGenerator{
			Target:    target,
			OutputDir: targetOutDir,
			Sessions:  sessions,
		}

		markdown := mdGen.Generate()
		mdPath := filepath.Join(targetOutDir, "comandos.md")
		if err := utils.WriteFile(mdPath, markdown); err != nil {
			utils.PrintError(fmt.Sprintf("Failed to write markdown for %s: %v", target, err))
			continue
		}

		// Generate plain text commands for this target
		txtGen := generator.PlainTextGenerator{
			Sessions: sessions,
		}

		plainText := txtGen.Generate()
		txtPath := filepath.Join(targetOutDir, "comandos.txt")
		if err := utils.WriteFile(txtPath, plainText); err != nil {
			utils.PrintError(fmt.Sprintf("Failed to write commands for %s: %v", target, err))
			continue
		}

		utils.PrintSuccess(fmt.Sprintf("Generated %d command(s) for %s", len(sessions), domain))
		utils.PrintInfo(fmt.Sprintf("  Markdown: %s", mdPath))
		utils.PrintInfo(fmt.Sprintf("  Commands: %s", txtPath))
		fmt.Println()
	}

	fmt.Println("📋 Quick Reference:")
	fmt.Printf("   Base output directory: %s\n", baseOutDir)
	fmt.Printf("   Domains file: %s\n", domainListFile)
	fmt.Printf("   Total targets: %d\n", len(targets))
	fmt.Println()
	utils.PrintInfo("Each target has its own subdirectory with comandos.md and comandos.txt")
	utils.PrintInfo("Review the files and execute manually or use 'trident-recon run'")

	return nil
}
