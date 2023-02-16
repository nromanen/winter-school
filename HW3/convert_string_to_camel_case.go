// https://www.codewars.com/kata/517abf86da9663f1d2000003

package kata

import (
	"strings"
	"unicode"
)

func ToCamelCase(s string) string {
	var result strings.Builder

	// Loop through each character in the string
	for i, c := range s {

		// If the character is a dash or underscore, skip it and capitalize the next character
		if c == '-' || c == '_' {
			continue
		}

		// If the previous character was a dash or underscore, capitalize the current character
		if i > 0 && (s[i-1] == '-' || s[i-1] == '_') {
			result.WriteRune(unicode.ToUpper(c))
		} else {
			result.WriteRune(c)
		}
	}

	return result.String()
}
