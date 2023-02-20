package kata

import (
	"strings"
)

//https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go

func ToCamelCase(s string) string {
	words := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(s, "-", " "), "_", " "))
	for i := 1; i < len(words); i++ {
		if words[i] != strings.Title(words[i]) {
			words[i] = strings.Title(words[i])
		}
	}
	return strings.Join(words, "")
}
