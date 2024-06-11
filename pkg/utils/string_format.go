package utils

import "strings"

// CapitalizeFirstLetter capitalizes the first letter of a string
func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
