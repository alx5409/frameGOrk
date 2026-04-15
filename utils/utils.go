package utils

import (
	"regexp"
	"strings"
)

// Compile the special characters on the module level for better performance, so that when is called again it does not
// compile again
var specialCharRE = regexp.MustCompile(`[^a-zA-Z0-9/_\-.]`)

// Detects whether the character is a special character
func specialCharacters(r rune) bool {
	return specialCharRE.MatchString(string(r))
}

// Normalizes the input method by removing spaces and special characters and converting to upper case
func NormalizeMethod(input string) string {
	output := strings.TrimFunc(input, specialCharacters) // Removes special characters
	return strings.ToUpper(output)
}

// Normalizes the path by removing the special characters
func NormalizePath(input string) string {
	return strings.TrimFunc(input, specialCharacters)
}
