package main

import (
	"fmt"
)

func convert(a, base int) (string, error) {
	if base < 2 || base > 36 {
		return "", fmt.Errorf("invalid base: %d", base)
	}
	if a == 0 {
		return "0", nil
	}
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := ""
	for a > 0 {
		digit := digits[a%base]
		result = string(digit) + result
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
