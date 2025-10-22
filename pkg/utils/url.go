package utils

import (
	"net/url"
	"strings"
)

// ParseURL parses a URL and extracts protocol and domain
func ParseURL(rawURL string) (protocol, domain string, err error) {
	// Add http:// if no protocol specified
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "http://" + rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", err
	}

	protocol = u.Scheme
	domain = u.Host

	return protocol, domain, nil
}

// SanitizeDomain removes port from domain if present
func SanitizeDomain(domain string) string {
	// Remove port if present
	if idx := strings.Index(domain, ":"); idx != -1 {
		domain = domain[:idx]
	}
	return domain
}

// NormalizeURL ensures URL has proper format
func NormalizeURL(rawURL string) string {
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		return "http://" + rawURL
	}
	return rawURL
}
