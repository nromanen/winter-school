package main

import (
	"encoding/json"
	"errors"
)

func convert(a, base int) (string, error) {
	if base < 2 || base > 36 {
		return "", errors.New("base must be between 2 and 36")
	}

	if a == 0 {
		return "0", nil
	}

	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := ""

	for a > 0 {
		index := a % base
		result = string(charset[index]) + result
		a /= base
	}

	return result, nil
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil && len(js) > 0
}

func Add(n int) func(int) int {
	return func(a int) int {
		return a + n
	}
}

func main() {
	conv, _ := convert(12, 16)
	println(conv)

	println(isJSON(`{"name": "John Doe", "age": 30}`))

	f := Add(5)
	println(f(10))
}
