package strings

import (
	"strings"
)

// RemoveExtension removes a known (.7z, .tar.gz, .rar...) extension of s and returns the result
// if the string doesn´t ends in the extension, returns the string unmodified
func RemoveExtension(s string) string {
	extensions := []string{".tar.gz", ".zip", ".7z", ".rar", ".tgz", ".tar", ".gz"}

	for _, ext := range extensions {
		if strings.HasSuffix(s, ext) {
			return strings.TrimSuffix(s, ext)
		}
	}

	return s
}

// RemoveString returns the s array without the r string
func RemoveString(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
