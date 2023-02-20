package kata

import (
	"sort"
	"strings"
)

// Task 3 https://www.codewars.com/kata/550554fd08b86f84fe000a58/train/go
func InArray(array1 []string, array2 []string) []string {
	joined := strings.Join(array2, " ")
	substrings := make(map[string]bool)
	for _, str := range array1 {
		if strings.Contains(joined, str) {
			substrings[str] = true
		}
	}
	var result = make([]string, 0, len(substrings))
	for k := range substrings {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}
