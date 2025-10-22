package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

// PromptOutputDir prompts the user for an output directory
func PromptOutputDir(defaultDir string) (string, error) {
	prompt := promptui.Prompt{
		Label:   "Output directory",
		Default: defaultDir,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	// Expand ~ to home directory
	if result[0] == '~' {
		home, _ := os.UserHomeDir()
		result = filepath.Join(home, result[1:])
	}

	return result, nil
}

// PromptConfirm prompts the user for confirmation
func PromptConfirm(label string) (bool, error) {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return false, err
	}

	return result == "y" || result == "Y", nil
}

// PromptSelect prompts the user to select from a list
func PromptSelect(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

// PrintSuccess prints a success message
func PrintSuccess(msg string) {
	fmt.Printf("✓ %s\n", msg)
}

// PrintError prints an error message
func PrintError(msg string) {
	fmt.Fprintf(os.Stderr, "✗ %s\n", msg)
}

// PrintInfo prints an info message
func PrintInfo(msg string) {
	fmt.Printf("ℹ %s\n", msg)
}

// PrintWarning prints a warning message
func PrintWarning(msg string) {
	fmt.Printf("⚠ %s\n", msg)
}
