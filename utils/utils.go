package utils

import "strings"

// Normalizes the input by removing spaces and converting to upper case
func Normalize(input string) string {
	ouput := strings.TrimSpace(input)
	return strings.ToUpper(ouput)
}
