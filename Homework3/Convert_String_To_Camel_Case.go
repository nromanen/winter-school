package main

import (
	"fmt"
	"strings"
)

func toCamelCase(s string) string {
	var result string
	value := false
	for i := 0; i < len(s); i++ {
		if s[i] == '-' || s[i] == '_' {
			value = true
		} else if value {
			result += strings.ToUpper(string(s[i]))
			value = false
		} else {
			result += string(s[i])
		}
	}
	return result
}

func main() {
	fmt.Printf("Input some text: ")
	var text string
	fmt.Scan(&text)
	fmt.Print(toCamelCase(text))
}
