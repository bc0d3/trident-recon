package generator

import (
	"fmt"
	"strings"
	"time"

	"github.com/bc0d3/trident-recon/pkg/executor"
)

// MarkdownGenerator generates markdown documentation
type MarkdownGenerator struct {
	Target    string
	OutputDir string
	Sessions  []executor.Session
}

// Generate generates the markdown content
func (mg *MarkdownGenerator) Generate() string {
	var md strings.Builder

	// Header
	md.WriteString("# 🔱 Trident Recon - Generated Commands\n\n")
	md.WriteString(fmt.Sprintf("**Target:** %s\n\n", mg.Target))
	md.WriteString(fmt.Sprintf("**Generated:** %s\n\n", time.Now().Format("2006-01-02 15:04:05")))
	md.WriteString(fmt.Sprintf("**Output Directory:** %s\n\n", mg.OutputDir))
	md.WriteString("---\n\n")

	// Quick Reference Table
	mg.generateQuickReference(&md)

	// Group sessions by tool
	toolSessions := mg.groupByTool()

	// Generate sections per tool
	for tool, sessions := range toolSessions {
		mg.generateToolSection(&md, tool, sessions)
	}

	// Session Management section
	mg.generateSessionManagement(&md)

	// Output Structure
	mg.generateOutputStructure(&md)

	// Results Analysis
	mg.generateResultsAnalysis(&md)

	return md.String()
}

func (mg *MarkdownGenerator) generateQuickReference(md *strings.Builder) {
	md.WriteString("## 📋 Quick Reference - Session IDs\n\n")
	md.WriteString("| Tool | Command | Session ID |\n")
	md.WriteString("|------|---------|------------|\n")
	for _, s := range mg.Sessions {
		md.WriteString(fmt.Sprintf("| %s | %s | `%s` |\n", s.Tool, s.CommandName, s.ID))
	}
	md.WriteString("\n---\n\n")
}

func (mg *MarkdownGenerator) groupByTool() map[string][]executor.Session {
	groups := make(map[string][]executor.Session)
	for _, s := range mg.Sessions {
		groups[s.Tool] = append(groups[s.Tool], s)
	}
	return groups
}

func (mg *MarkdownGenerator) generateToolSection(md *strings.Builder, tool string, sessions []executor.Session) {
	emoji := mg.getToolEmoji(tool)
	md.WriteString(fmt.Sprintf("## %s %s Commands\n\n", emoji, strings.ToUpper(tool)))

	for i, s := range sessions {
		md.WriteString(fmt.Sprintf("### %d. %s\n\n", i+1, s.CommandName))
		md.WriteString(fmt.Sprintf("**Session ID:** `%s`\n\n", s.ID))
		md.WriteString(fmt.Sprintf("**Description:** %s\n\n", s.CommandName))
		if s.Wordlist != "" {
			md.WriteString(fmt.Sprintf("**Wordlist:** `%s`\n\n", s.Wordlist))
		}

		md.WriteString("```bash\n")
		md.WriteString("# Start session\n")
		md.WriteString(fmt.Sprintf("tmux new-session -d -s \"%s\" bash -c \"%s\"\n\n", s.TmuxSession, escapeCommand(s.Command)))
		md.WriteString("# Attach to session\n")
		md.WriteString(fmt.Sprintf("tmux attach -t \"%s\"\n\n", s.TmuxSession))
		if s.OutputFile != "" {
			md.WriteString("# View output\n")
			md.WriteString(fmt.Sprintf("cat %s\n", s.OutputFile))
		}
		md.WriteString("```\n\n")
		md.WriteString("---\n\n")
	}
}

func (mg *MarkdownGenerator) getToolEmoji(tool string) string {
	emojis := map[string]string{
		"ffuf":        "🎯",
		"gobuster":    "🔍",
		"dirsearch":   "🌐",
		"feroxbuster": "⚡",
		"httpx":       "🌍",
		"nuclei":      "💣",
		"arjun":       "🔑",
		"katana":      "⚔️",
		"waybackurls": "📜",
		"gau":         "🗂️",
		"gospider":    "🕷️",
		"hakrawler":   "🦀",
		"dalfox":      "🎭",
		"sqlmap":      "💉",
	}
	if emoji, ok := emojis[tool]; ok {
		return emoji
	}
	return "🔧"
}

func (mg *MarkdownGenerator) generateSessionManagement(md *strings.Builder) {
	md.WriteString("## 🎮 Session Management\n\n")
	md.WriteString("### List all active Trident sessions\n\n")
	md.WriteString("```bash\n")
	md.WriteString("# Using trident-recon\n")
	md.WriteString("trident-recon list\n\n")
	md.WriteString("# Using tmux directly\n")
	md.WriteString("tmux ls | grep -E \"(ffuf_|gobuster_|dirsearch_|feroxbuster_|httpx_|nuclei_)\"\n")
	md.WriteString("```\n\n")

	md.WriteString("### Attach to a session\n\n")
	md.WriteString("```bash\n")
	md.WriteString("tmux attach -t <session-name>\n")
	md.WriteString("# Detach: Ctrl+B then D\n")
	md.WriteString("```\n\n")

	md.WriteString("### Kill sessions\n\n")
	md.WriteString("```bash\n")
	md.WriteString("# Kill specific session\n")
	md.WriteString("trident-recon kill <session-id>\n\n")
	md.WriteString("# Kill all trident sessions\n")
	md.WriteString("trident-recon kill-all\n\n")
	md.WriteString("# Kill all sessions for a specific tool\n")
	md.WriteString("trident-recon kill-all --tool ffuf\n")
	md.WriteString("```\n\n")
	md.WriteString("---\n\n")
}

func (mg *MarkdownGenerator) generateOutputStructure(md *strings.Builder) {
	md.WriteString("## 📁 Output Structure\n\n")
	md.WriteString("```\n")
	md.WriteString(fmt.Sprintf("%s/\n", mg.OutputDir))

	toolGroups := mg.groupByTool()
	for tool, sessions := range toolGroups {
		md.WriteString(fmt.Sprintf("├── %s/\n", tool))
		for _, s := range sessions {
			if s.OutputFile != "" {
				md.WriteString(fmt.Sprintf("│   ├── %s\n", extractFileName(s.OutputFile)))
			}
		}
	}
	md.WriteString("└── commands.md (this file)\n")
	md.WriteString("```\n\n")
	md.WriteString("---\n\n")
}

func (mg *MarkdownGenerator) generateResultsAnalysis(md *strings.Builder) {
	md.WriteString("## 📊 Results Analysis\n\n")
	md.WriteString("### View all JSON results with jq\n\n")
	md.WriteString("```bash\n")
	md.WriteString(fmt.Sprintf("cd %s\n\n", mg.OutputDir))
	md.WriteString("# View ffuf results\n")
	md.WriteString("for file in ffuf-*.json; do\n")
	md.WriteString("    echo \"=== $file ===\"\n")
	md.WriteString("    jq '.results[] | {url: .url, status: .status, length: .length}' \"$file\" | head -20\n")
	md.WriteString("done\n")
	md.WriteString("```\n\n")

	md.WriteString("### Check running sessions\n\n")
	md.WriteString("```bash\n")
	md.WriteString("# List all active sessions\n")
	md.WriteString("trident-recon list\n\n")
	md.WriteString("# Monitor a specific session\n")
	md.WriteString("tmux attach -t <session-name>\n")
	md.WriteString("```\n\n")

	md.WriteString("---\n\n")
	md.WriteString("## 🎯 Next Steps\n\n")
	md.WriteString("1. Review the generated commands above\n")
	md.WriteString("2. Execute commands using `trident-recon run` or manually\n")
	md.WriteString("3. Monitor sessions with `tmux attach`\n")
	md.WriteString("4. Analyze results in the output directory\n")
	md.WriteString("5. Kill sessions when done with `trident-recon kill-all`\n\n")
}

func escapeCommand(cmd string) string {
	// Escape double quotes for bash -c
	return strings.ReplaceAll(cmd, `"`, `\"`)
}

func extractFileName(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return path
}
