package main

import (
    "fmt"
    "regexp"
)

//Task 2
func isJSON(str string) bool {
    re := regexp.MustCompile(`^\s*\{\s*"([^"]*)"\s*:\s*("[^"]*"|\d*|\{[^{}]*\})(,\s*"([^"]*)"\s*:\s*("[^"]*"|\d*|\{[^{}]*\}))*\s*\}\s*$`)
    return re.MatchString(str)
}

func main() {
    // Examples of JSON strings that match the format
    json1 := `{"name": "John", "age": 30, "city": {"name": "New York", "population": 100000}}`
    json2 := `{"name": "Alice", "age": 25}`
    
    // Example of non-matching string
    json3 := `{"name": "Bob", "age": 35, "cities": ["New York", "Los Angeles"]}`
    
    // Test the strings with the isJSON function
    fmt.Println(isJSON(json1)) // true
    fmt.Println(isJSON(json2)) // true
    fmt.Println(isJSON(json3)) // false
}