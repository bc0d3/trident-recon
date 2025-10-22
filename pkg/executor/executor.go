package executor

import (
	"fmt"
	"time"

	"github.com/bc0d3/trident-recon/pkg/tmux"
	"github.com/bc0d3/trident-recon/pkg/utils"
)

// Executor executes commands in tmux sessions
type Executor struct {
	StateDir string
}

// NewExecutor creates a new executor
func NewExecutor(stateDir string) *Executor {
	return &Executor{
		StateDir: stateDir,
	}
}

// Execute executes a single session
func (e *Executor) Execute(session *Session) error {
	// Ensure output directory exists
	if err := utils.EnsureDir(session.OutputDir); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Check if tmux is available
	if !tmux.IsTmuxAvailable() {
		return fmt.Errorf("tmux is not installed or not available")
	}

	// Check if session already exists
	if tmux.SessionExists(session.TmuxSession) {
		return fmt.Errorf("tmux session %s already exists", session.TmuxSession)
	}

	// Set started time
	session.StartedAt = time.Now()
	session.Status = "running"

	// Create tmux session
	if err := tmux.CreateSession(session.TmuxSession, session.Command); err != nil {
		return fmt.Errorf("failed to create tmux session: %w", err)
	}

	// Save session metadata
	if err := session.Save(e.StateDir); err != nil {
		// Try to cleanup tmux session if metadata save fails
		tmux.KillSession(session.TmuxSession)
		return fmt.Errorf("failed to save session metadata: %w", err)
	}

	return nil
}

// ExecuteAll executes multiple sessions
func (e *Executor) ExecuteAll(sessions []Session) (int, error) {
	successful := 0
	failed := 0

	for i, session := range sessions {
		utils.PrintInfo(fmt.Sprintf("[%d/%d] Executing %s - %s...", i+1, len(sessions), session.Tool, session.CommandName))

		if err := e.Execute(&session); err != nil {
			utils.PrintError(fmt.Sprintf("Failed to execute %s - %s: %v", session.Tool, session.CommandName, err))
			failed++
			continue
		}

		utils.PrintSuccess(fmt.Sprintf("Started session %s (ID: %s)", session.TmuxSession, session.ID))
		successful++
	}

	return successful, nil
}

// ValidateSessions validates that all sessions can be executed
func (e *Executor) ValidateSessions(sessions []Session) error {
	// Check if tmux is available
	if !tmux.IsTmuxAvailable() {
		return fmt.Errorf("tmux is not installed or not available")
	}

	// Check for session name conflicts
	existingSessions, err := tmux.ListSessions()
	if err != nil {
		existingSessions = []string{}
	}

	conflicts := make(map[string]bool)
	for _, existing := range existingSessions {
		conflicts[existing] = true
	}

	for _, session := range sessions {
		if conflicts[session.TmuxSession] {
			return fmt.Errorf("session %s already exists in tmux", session.TmuxSession)
		}
	}

	return nil
}
