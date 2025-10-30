package generator

import (
	"fmt"
	"os"

	"github.com/bc0d3/trident-recon/pkg/config"
	"github.com/bc0d3/trident-recon/pkg/executor"
	"github.com/bc0d3/trident-recon/pkg/utils"
)

// Generator handles command generation
type Generator struct {
	Config         *config.Config
	Target         string
	OutputDir      string
	DomainListFile string
}

// New creates a new generator
func New(cfg *config.Config, target, outputDir string) *Generator {
	return &Generator{
		Config:    cfg,
		Target:    target,
		OutputDir: outputDir,
	}
}

// SetDomainListFile sets the domain list file path
func (g *Generator) SetDomainListFile(path string) {
	g.DomainListFile = path
}

// Generate generates all commands for enabled tools
func (g *Generator) Generate(toolsFilter, skipTools []string) ([]executor.Session, error) {
	protocol, domain, err := utils.ParseURL(g.Target)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	domain = utils.SanitizeDomain(domain)
	normalizedURL := utils.NormalizeURL(g.Target)

	var sessions []executor.Session

	// Iterate through enabled tools
	for toolName, toolConfig := range g.Config.Tools {
		if !toolConfig.Enabled {
			continue
		}

		// Apply filters
		if len(toolsFilter) > 0 && !contains(toolsFilter, toolName) {
			continue
		}
		if len(skipTools) > 0 && contains(skipTools, toolName) {
			continue
		}

		// Generate commands for this tool
		for _, cmdTemplate := range toolConfig.Commands {
			session := g.generateSession(toolName, toolConfig, cmdTemplate, normalizedURL, protocol, domain)
			sessions = append(sessions, session)
		}
	}

	return sessions, nil
}

func (g *Generator) generateSession(toolName string, toolConfig config.ToolConfig, cmdTemplate config.CommandTemplate, url, protocol, domain string) executor.Session {
	// Generate unique ID
	id := utils.GenerateID(toolName, cmdTemplate.Name, domain)

	// Get wordlist path if specified
	wordlist := ""
	if cmdTemplate.Wordlist != "" {
		if path, ok := g.Config.Wordlists[cmdTemplate.Wordlist]; ok {
			wordlist = os.ExpandEnv(path)
		}
	}

	// Build dynamic headers map from config
	headersMap := BuildHeadersMap(g.Config.Headers)

	// Create replacements
	replacements := Replacements{
		URL:        url,
		Domain:     domain,
		Protocol:   protocol,
		Wordlist:   wordlist,
		OutputDir:  g.OutputDir,
		ID:         id,
		DomainList: g.DomainListFile,
		Headers:    headersMap,
	}

	// Replace template variables
	command := ReplaceTemplateVars(cmdTemplate.Command, replacements)

	// Generate tmux session name
	tmuxSession := fmt.Sprintf("%s%s", toolConfig.TmuxPrefix, id)

	// Determine output file from command if possible
	outputFile := ""
	if strings := findOutputFlag(command); strings != "" {
		outputFile = strings
	}

	return executor.Session{
		ID:          id,
		Tool:        toolName,
		CommandName: cmdTemplate.Name,
		Target:      url,
		TmuxSession: tmuxSession,
		Command:     command,
		OutputDir:   g.OutputDir,
		OutputFile:  outputFile,
		Wordlist:    wordlist,
		Status:      "pending",
	}
}

// findOutputFlag tries to extract output file path from command
func findOutputFlag(command string) string {
	// Simple extraction of output file paths
	// Look for common patterns like -o file, --output file, > file
	parts := splitCommand(command)
	for i, part := range parts {
		if (part == "-o" || part == "--output" || part == "-oJ") && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

// splitCommand splits a command string into parts (simple implementation)
func splitCommand(cmd string) []string {
	var parts []string
	var current string
	inQuote := false

	for _, c := range cmd {
		if c == '"' || c == '\'' {
			inQuote = !inQuote
		} else if c == ' ' && !inQuote {
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
		} else {
			current += string(c)
		}
	}

	if current != "" {
		parts = append(parts, current)
	}

	return parts
}

// contains checks if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
