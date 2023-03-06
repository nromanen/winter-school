package main

import (
	"errors"
	"fmt"
)

func convert(a, base int) (string, error) {
	if base < 2 || base > 36 {
		return "", errors.New("invalid base")
	}
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if a == 0 {
		return "0", nil
	}
	result := ""
	for a > 0 {
		digit := digits[a%base]
		result = string(digit) + result
		a /= base
	}
	return result, nil
}

func main() {

	result, err := convert(10, 16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

}
