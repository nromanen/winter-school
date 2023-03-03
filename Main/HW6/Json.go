package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func main() {
	// Example JSON strings
	jsonString1 := `{"name": "John", "age": 30}`
	jsonString2 := `{"name": "John", "age": 30,}` // invalid JSON string

	// Regular expression to match JSON without arrays
	jsonPattern := `^{\s*("[\w]+")\s*:\s*("[^"]*"|\d*)(,\s*("[\w]+")\s*:\s*("[^"]*"|\d*))*\s*}$`

	// Check if the string matches the pattern
	match, _ := regexp.MatchString(jsonPattern, jsonString1)
	if match {
		// Parse the JSON string
		var data interface{}
		err := json.Unmarshal([]byte(jsonString1), &data)
		if err != nil {
			fmt.Println("Invalid JSON")
			return
		}
		fmt.Println("Valid JSON")
	} else {
		fmt.Println("Invalid JSON")
	}

	// Check if the second string matches the pattern
	match, _ = regexp.MatchString(jsonPattern, jsonString2)
	if match {
		// Parse the JSON string
		var data interface{}
		err := json.Unmarshal([]byte(jsonString2), &data)
		if err != nil {
			fmt.Println("Invalid JSON")
			return
		}
		fmt.Println("Valid JSON")
	} else {
		fmt.Println("Invalid JSON")
	}
}