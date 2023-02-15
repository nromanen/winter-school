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
	// expected - "theStealthWarrior"
	fmt.Println(ToCamelCase("the-stealth-warrior"))
	// expected - "TheStealthWarrior"
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
}
