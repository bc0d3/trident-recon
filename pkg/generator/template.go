package generator

import (
	"fmt"
	"strings"

	"github.com/bc0d3/trident-recon/pkg/config"
)

// Replacements holds template variable replacements
type Replacements struct {
	URL        string
	Domain     string
	Protocol   string
	Wordlist   string
	OutputDir  string
	ID         string
	DomainList string
	Headers    map[string]string // Dynamic header replacements
}

// ReplaceTemplateVars replaces template variables in a command string
func ReplaceTemplateVars(template string, rep Replacements) string {
	result := template

	// Replace basic variables
	result = strings.ReplaceAll(result, "{URL}", rep.URL)
	result = strings.ReplaceAll(result, "{DOMAIN}", rep.Domain)
	result = strings.ReplaceAll(result, "{PROTOCOL}", rep.Protocol)
	result = strings.ReplaceAll(result, "{WORDLIST}", rep.Wordlist)
	result = strings.ReplaceAll(result, "{OUTPUT_DIR}", rep.OutputDir)
	result = strings.ReplaceAll(result, "{ID}", rep.ID)
	result = strings.ReplaceAll(result, "{DOMAIN_LIST}", rep.DomainList)

	// Replace all header variables dynamically
	// This supports {HEADER-User-Agent}, {HEADER-X-Bug-Bounty}, {HEADERS-ALL}, etc.
	for placeholder, value := range rep.Headers {
		result = strings.ReplaceAll(result, placeholder, value)
	}

	return result
}

// BuildHeadersMap creates a map of ALL possible header replacement variables
// This function dynamically creates placeholders for ANY header in the config
//
// Examples:
//   {HEADER-User-Agent} → -H "User-Agent: Mozilla/5.0..."
//   {HEADER-Accept} → -H "Accept: */*"
//   {HEADER-X-Bug-Bounty} → -H "X-Bug-Bounty: hackeroneUser"
//   {HEADERS-ALL} → all headers combined
//   {HEADERS-DEFAULT} → only default headers
//   {HEADERS-CUSTOM} → only custom headers
func BuildHeadersMap(headers config.HeadersConfig) map[string]string {
	result := make(map[string]string)

	var defaultParts []string
	var customParts []string

	// Process ALL default headers dynamically
	for key, value := range headers.Default {
		formatted := fmt.Sprintf(`-H "%s: %s"`, key, value)

		// Create placeholder: {HEADER-User-Agent}, {HEADER-Accept}, etc.
		placeholder := fmt.Sprintf("{HEADER-%s}", key)
		result[placeholder] = formatted

		defaultParts = append(defaultParts, formatted)
	}

	// Process ALL custom headers dynamically
	for _, header := range headers.Custom {
		// Parse custom header to extract the key name
		// Supports both: "X-Bug-Bounty: value" and "X-Custom-Header"
		parts := strings.SplitN(header, ":", 2)
		headerKey := strings.TrimSpace(parts[0])

		formatted := fmt.Sprintf(`-H "%s"`, header)

		// Create placeholder: {HEADER-X-Bug-Bounty}, {HEADER-X-Custom}, etc.
		placeholder := fmt.Sprintf("{HEADER-%s}", headerKey)
		result[placeholder] = formatted

		customParts = append(customParts, formatted)
	}

	// Build grouped placeholders
	result["{HEADERS-DEFAULT}"] = strings.Join(defaultParts, " ")
	result["{HEADERS-CUSTOM}"] = strings.Join(customParts, " ")

	// Combine all headers (default + custom)
	allParts := append([]string{}, defaultParts...)
	allParts = append(allParts, customParts...)
	result["{HEADERS-ALL}"] = strings.Join(allParts, " ")

	return result
}
