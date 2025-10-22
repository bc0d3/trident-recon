package generator

import (
	"fmt"
	"strings"

	"github.com/bc0d3/trident-recon/pkg/executor"
)

// PlainTextGenerator generates plain text commands (copy-paste ready)
type PlainTextGenerator struct {
	Sessions []executor.Session
}

// Generate generates plain text commands without any formatting
func (tg *PlainTextGenerator) Generate() string {
	var txt strings.Builder

	for _, s := range tg.Sessions {
		// Just write the raw tmux command that can be copy-pasted
		txt.WriteString(fmt.Sprintf("tmux new-session -d -s \"%s\" bash -c \"%s\"\n",
			s.TmuxSession,
			escapeForBash(s.Command)))
	}

	return txt.String()
}

// escapeForBash escapes special characters for bash -c execution
func escapeForBash(cmd string) string {
	// Escape double quotes and backslashes for bash -c
	cmd = strings.ReplaceAll(cmd, `\`, `\\`)
	cmd = strings.ReplaceAll(cmd, `"`, `\"`)
	cmd = strings.ReplaceAll(cmd, `$`, `\$`)
	return cmd
}
