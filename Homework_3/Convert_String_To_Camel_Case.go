package kata

import (
	"strings"
)

func ToCamelCase(s string) string {
	result := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(s, "-", " "), "_", " "))
	for i := 1; i < len(result); i++ {
		if result[i] != strings.Title(result[i]) {
			result[i] = strings.Title(result[i])
		}
	}

	return strings.Join(result, "")
}
