package main

import (
	"fmt"
	"strings"
)

func ToCamelCase(s string) string {
	words := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(s, "-", " "), "_", " "))
	for i := 1; i < len(words); i++ {
		if words[i] != strings.Title(words[i]) {
			words[i] = strings.Title(words[i])
		}
	}
	return strings.Join(words, "")
}

func main() {
	fmt.Println(ToCamelCase("dark-souls-two"))
	fmt.Println(ToCamelCase("Dark_Souls_Two"))
}
