package kata

import (
	"strings"
)

// Task 3 https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go
func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}
	result := strings.Title(strings.Replace(strings.Replace(s, "-", " ", -1), "_", " ", -1))
	result = s[:1] + result[1:]
	result = strings.Replace(result, " ", "", -1)
	return result
}
