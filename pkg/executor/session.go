package executor

import (
	"fmt"

	"github.com/bc0d3/trident-recon/pkg/tmux"
)

// SessionManager manages tmux sessions
type SessionManager struct {
	StateDir string
}

// NewSessionManager creates a new session manager
func NewSessionManager(stateDir string) *SessionManager {
	return &SessionManager{
		StateDir: stateDir,
	}
}

// ListSessions lists all active sessions
func (sm *SessionManager) ListSessions(toolFilter string) ([]Session, error) {
	// Load all saved sessions
	sessions, err := LoadAll(sm.StateDir)
	if err != nil {
		return nil, err
	}

	// Get active tmux sessions
	tmuxSessions, err := tmux.ListSessions()
	if err != nil {
		// If no tmux sessions, return empty list
		tmuxSessions = []string{}
	}

	// Create map of active tmux sessions
	activeSessions := make(map[string]bool)
	for _, ts := range tmuxSessions {
		activeSessions[ts] = true
	}

	// Filter sessions
	var activeTridentSessions []Session
	for _, s := range sessions {
		// Check if session is still active in tmux
		if !activeSessions[s.TmuxSession] {
			s.Status = "completed"
		} else {
			s.Status = "running"
		}

		// Apply tool filter
		if toolFilter != "" && s.Tool != toolFilter {
			continue
		}

		activeTridentSessions = append(activeTridentSessions, s)
	}

	return activeTridentSessions, nil
}

// KillSession kills a specific session
func (sm *SessionManager) KillSession(id string) error {
	// Load session metadata
	session, err := Load(sm.StateDir, id)
	if err != nil {
		return fmt.Errorf("session not found: %w", err)
	}

	// Kill tmux session
	if tmux.SessionExists(session.TmuxSession) {
		if err := tmux.KillSession(session.TmuxSession); err != nil {
			return fmt.Errorf("failed to kill tmux session: %w", err)
		}
	}

	// Delete metadata
	if err := Delete(sm.StateDir, id); err != nil {
		return fmt.Errorf("failed to delete session metadata: %w", err)
	}

	return nil
}

// KillAllSessions kills all sessions, optionally filtered by tool
func (sm *SessionManager) KillAllSessions(toolFilter string) (int, error) {
	sessions, err := sm.ListSessions(toolFilter)
	if err != nil {
		return 0, err
	}

	killed := 0
	for _, session := range sessions {
		if err := sm.KillSession(session.ID); err != nil {
			fmt.Printf("Warning: failed to kill session %s: %v\n", session.ID, err)
			continue
		}
		killed++
	}

	return killed, nil
}

// GetSession gets a specific session by ID
func (sm *SessionManager) GetSession(id string) (*Session, error) {
	return Load(sm.StateDir, id)
}

// AttachToSession attaches to a tmux session
func (sm *SessionManager) AttachToSession(id string) error {
	session, err := Load(sm.StateDir, id)
	if err != nil {
		return fmt.Errorf("session not found: %w", err)
	}

	if !tmux.SessionExists(session.TmuxSession) {
		return fmt.Errorf("tmux session no longer exists")
	}

	return tmux.AttachSession(session.TmuxSession)
}
