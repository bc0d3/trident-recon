package executor

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Session represents a command execution session
type Session struct {
	ID          string    `json:"id"`
	Tool        string    `json:"tool"`
	CommandName string    `json:"command_name"`
	Target      string    `json:"target"`
	TmuxSession string    `json:"tmux_session"`
	Command     string    `json:"command"`
	OutputDir   string    `json:"output_dir"`
	OutputFile  string    `json:"output_file"`
	Wordlist    string    `json:"wordlist"`
	StartedAt   time.Time `json:"started_at"`
	Status      string    `json:"status"`
}

// Save saves session metadata to disk
func (s *Session) Save(stateDir string) error {
	jobsDir := filepath.Join(stateDir, "jobs")
	if err := os.MkdirAll(jobsDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	filename := filepath.Join(jobsDir, s.ID+".json")
	return os.WriteFile(filename, data, 0644)
}

// Load loads a session from disk
func Load(stateDir, id string) (*Session, error) {
	filename := filepath.Join(stateDir, "jobs", id+".json")
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var session Session
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

// LoadAll loads all sessions from disk
func LoadAll(stateDir string) ([]Session, error) {
	jobsDir := filepath.Join(stateDir, "jobs")

	// Check if jobs directory exists
	if _, err := os.Stat(jobsDir); os.IsNotExist(err) {
		return []Session{}, nil
	}

	entries, err := os.ReadDir(jobsDir)
	if err != nil {
		return nil, err
	}

	var sessions []Session
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(jobsDir, entry.Name()))
		if err != nil {
			continue
		}

		var session Session
		if err := json.Unmarshal(data, &session); err != nil {
			continue
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
}

// Delete deletes a session from disk
func Delete(stateDir, id string) error {
	filename := filepath.Join(stateDir, "jobs", id+".json")
	return os.Remove(filename)
}
