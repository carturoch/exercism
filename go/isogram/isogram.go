package isogram

import (
	"strings"
)

// IsIsogram determines whether a word is composed only
// by unique characters
func IsIsogram(s string) bool {
	if s == "" {
		return true
	}
	iso := ""
	for _, c := range strings.ToLower(s) {
		if c == '-' || c == ' ' {
			continue
		}
		if strings.ContainsRune(iso, rune(c)) {
			return false
		}
		iso += string(c)
	}
	return true
}
