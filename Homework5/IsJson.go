package main

import (
	"encoding/json"
	"fmt"
)

func isValidJSON(str string) bool {
	var obj map[string]interface{}

	if err := json.Unmarshal([]byte(str), &obj); err != nil {
		return false
	}

	return true
}

func main() {
	// приклад рядків, які перевіряємо
	jsonStr1 := `{"name": "John Doe", "age": 30, "isMarried": false}`
	jsonStr2 := `Trello, aloo llll`

	// перевірка, чи є рядки json-форматом без масивів
	if isValidJSON(jsonStr1) {
		fmt.Println("Valid JSON:", jsonStr1)
	} else {
		fmt.Println("Invalid JSON format:", jsonStr1)
	}

	if isValidJSON(jsonStr2) {
		fmt.Println("Valid JSON:", jsonStr2)
	} else {
		fmt.Println("Invalid JSON format:", jsonStr2)
	}
}
