package config

import (
	"fmt"
	"os"
)

// Validate validates the configuration
func (c *Config) Validate() error {
	// Validate global config
	if c.Global.OutputDir == "" {
		return fmt.Errorf("global.output_dir cannot be empty")
	}

	if c.Global.IDLength <= 0 {
		c.Global.IDLength = 12 // default value
	}

	// Validate wordlists existence (warn only)
	for name, path := range c.Wordlists {
		expandedPath := os.ExpandEnv(path)
		if _, err := os.Stat(expandedPath); os.IsNotExist(err) {
			fmt.Printf("Warning: wordlist '%s' not found at %s\n", name, expandedPath)
		}
	}

	// Validate tools
	if len(c.Tools) == 0 {
		return fmt.Errorf("no tools configured")
	}

	for toolName, tool := range c.Tools {
		if tool.Enabled {
			if tool.TmuxPrefix == "" {
				return fmt.Errorf("tool %s: tmux_prefix cannot be empty", toolName)
			}
			if len(tool.Commands) == 0 {
				return fmt.Errorf("tool %s: no commands configured", toolName)
			}
			for i, cmd := range tool.Commands {
				if cmd.Name == "" {
					return fmt.Errorf("tool %s: command %d has no name", toolName, i)
				}
				if cmd.Command == "" {
					return fmt.Errorf("tool %s: command %s has no command template", toolName, cmd.Name)
				}
			}
		}
	}

	return nil
}

// GetEnabledTools returns a list of enabled tool names
func (c *Config) GetEnabledTools() []string {
	var enabled []string
	for name, tool := range c.Tools {
		if tool.Enabled {
			enabled = append(enabled, name)
		}
	}
	return enabled
}

// GetToolConfig returns the config for a specific tool
func (c *Config) GetToolConfig(toolName string) (*ToolConfig, error) {
	tool, ok := c.Tools[toolName]
	if !ok {
		return nil, fmt.Errorf("tool %s not found in config", toolName)
	}
	return &tool, nil
}
