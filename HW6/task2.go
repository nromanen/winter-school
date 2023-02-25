package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func main() {
	jsonString := `{"name": "John Doe", "age": 30}`

	// Define regular expression to match JSON object without arrays
	jsonRegexp := regexp.MustCompile(`^\s*\{(\s*"[^"]+"\s*:\s*(?:(?:"[^"]*")|(?:-?\d+(?:\.\d+)?(?:[eE][-+]?\d+)?))\s*(?:,\s*(?:"[^"]+"\s*:\s*(?:(?:"[^"]*")|(?:-?\d+(?:\.\d+)?(?:[eE][-+]?\d+)?))\s*))*)?\}\s*$`)

	// Check if the string matches the regular expression
	if !jsonRegexp.MatchString(jsonString) {
		fmt.Println("The string is not valid JSON.")
		return
	}

	// Check if the string is valid JSON using json.Valid
	if !json.Valid([]byte(jsonString)) {
		fmt.Println("The string is not valid JSON.")
		return
	}

	fmt.Println("The string is valid JSON.")
}
