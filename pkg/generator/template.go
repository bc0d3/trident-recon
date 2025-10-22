package generator

import (
	"strings"
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
}

// ReplaceTemplateVars replaces template variables in a command string
func ReplaceTemplateVars(template string, rep Replacements) string {
	result := template
	result = strings.ReplaceAll(result, "{URL}", rep.URL)
	result = strings.ReplaceAll(result, "{DOMAIN}", rep.Domain)
	result = strings.ReplaceAll(result, "{PROTOCOL}", rep.Protocol)
	result = strings.ReplaceAll(result, "{WORDLIST}", rep.Wordlist)
	result = strings.ReplaceAll(result, "{OUTPUT_DIR}", rep.OutputDir)
	result = strings.ReplaceAll(result, "{ID}", rep.ID)
	result = strings.ReplaceAll(result, "{DOMAIN_LIST}", rep.DomainList)
	return result
}
