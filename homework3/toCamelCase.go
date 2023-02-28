package main

import (
	"fmt"
	"strings"
)

func toCamelCase(s string) string {
	var camelCase string
	var capitalizeNext bool
	for i := 0; i < len(s); i++ {
		if s[i] == '-' || s[i] == '_' {
			capitalizeNext = true
		} else if capitalizeNext {
			camelCase += strings.ToUpper(string(s[i]))
			capitalizeNext = false
		} else {
			camelCase += string(s[i])
		}
	}
	return camelCase
}

func main() {

	fmt.Println(toCamelCase("the-stealth-warrior")) // prints "theStealthWarrior"
	fmt.Println(toCamelCase("The_Stealth_Warrior")) // prints "TheStealthWarrior"

}
