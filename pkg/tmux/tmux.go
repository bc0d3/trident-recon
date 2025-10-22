package tmux

import (
	"fmt"
	"os/exec"
	"strings"
)

// CreateSession creates a new tmux session
func CreateSession(sessionName, command string) error {
	cmd := exec.Command("tmux", "new-session", "-d", "-s", sessionName, "bash", "-c", command)
	return cmd.Run()
}

// SessionExists checks if a tmux session exists
func SessionExists(sessionName string) bool {
	cmd := exec.Command("tmux", "has-session", "-t", sessionName)
	return cmd.Run() == nil
}

// ListSessions lists all tmux sessions
func ListSessions() ([]string, error) {
	cmd := exec.Command("tmux", "list-sessions", "-F", "#{session_name}")
	output, err := cmd.Output()
	if err != nil {
		// If tmux is not running or no sessions exist, return empty list
		if strings.Contains(err.Error(), "no server running") {
			return []string{}, nil
		}
		return nil, err
	}

	if len(output) == 0 {
		return []string{}, nil
	}

	sessions := strings.Split(strings.TrimSpace(string(output)), "\n")
	return sessions, nil
}

// KillSession kills a tmux session
func KillSession(sessionName string) error {
	cmd := exec.Command("tmux", "kill-session", "-t", sessionName)
	return cmd.Run()
}

// AttachSession attaches to a tmux session
func AttachSession(sessionName string) error {
	cmd := exec.Command("tmux", "attach-session", "-t", sessionName)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

// FilterSessionsByPrefix filters sessions by prefix
func FilterSessionsByPrefix(sessions []string, prefix string) []string {
	var filtered []string
	for _, s := range sessions {
		if strings.HasPrefix(s, prefix) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

// GetSessionInfo returns information about a session
func GetSessionInfo(sessionName string) (map[string]string, error) {
	if !SessionExists(sessionName) {
		return nil, fmt.Errorf("session %s does not exist", sessionName)
	}

	cmd := exec.Command("tmux", "display-message", "-t", sessionName, "-p", "#{session_name}:#{session_created}:#{session_windows}")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	parts := strings.Split(strings.TrimSpace(string(output)), ":")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid session info format")
	}

	info := map[string]string{
		"name":    parts[0],
		"created": parts[1],
		"windows": parts[2],
	}

	return info, nil
}

// IsTmuxAvailable checks if tmux is installed
func IsTmuxAvailable() bool {
	cmd := exec.Command("tmux", "-V")
	return cmd.Run() == nil
}
