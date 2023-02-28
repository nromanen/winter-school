package main

//Перевести задане десяткове число в задану систему числення convert(a, base int) (string, error).

//convert(12, 16) - > "C"

import (
	"errors"
	"fmt"
)

func convert(a, base int) (string, error) {
	if base < 2 || base > 36 {
		return "", errors.New("base must be between 2 and 36")
	}
	if a < 0 {
		return "", errors.New("a must be a non-negative integer")
	}
	if a == 0 {
		return "0", nil
	}
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result string
	for a > 0 {
		remainder := a % base
		result = string(digits[remainder]) + result
		a = a / base
	}
	return result, nil
}

func main() {
	result, err := convert(12, 16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
