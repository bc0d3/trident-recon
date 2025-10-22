package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents the main configuration structure
type Config struct {
	Global    GlobalConfig          `yaml:"global"`
	Headers   HeadersConfig         `yaml:"headers"`
	Tools     map[string]ToolConfig `yaml:"tools"`
	Wordlists map[string]string     `yaml:"wordlists"`
}

// GlobalConfig contains global settings
type GlobalConfig struct {
	OutputDir string `yaml:"output_dir"`
	IDLength  int    `yaml:"id_length"`
}

// HeadersConfig contains HTTP headers configuration
type HeadersConfig struct {
	Default map[string]string `yaml:"default"`
	Custom  []string          `yaml:"custom"`
}

// ToolConfig represents a tool configuration
type ToolConfig struct {
	Enabled    bool              `yaml:"enabled"`
	TmuxPrefix string            `yaml:"tmux_prefix"`
	Commands   []CommandTemplate `yaml:"commands"`
}

// CommandTemplate represents a command template
type CommandTemplate struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Command     string `yaml:"command"`
	Wordlist    string `yaml:"wordlist"`
}

// Load reads and parses the config file
func Load() (*Config, error) {
	home, _ := os.UserHomeDir()
	configPath := filepath.Join(home, ".config", "trident-recon", "config.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("config not found at %s: %w", configPath, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}

	return &cfg, nil
}

// GetConfigPath returns the config file path
func GetConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "trident-recon", "config.yaml")
}

// GetStateDir returns the state directory path
func GetStateDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".local", "state", "trident-recon")
}
