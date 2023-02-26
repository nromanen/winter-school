package main

import (
	"errors"
	"fmt"
)

//Task 1
func convert(a, base int) (string, error) {
	if base < 2 || base > 36 {
		return "", errors.New("base must be between 2 and 36")
	}
	if a == 0 {
		return "0", nil
	}
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := ""
	for a > 0 {
		digit := digits[a%base]
		result = string(digit) + result
		a /= base
	}
	return result, nil
}

func main() {
	// Example: Convert decimal 42 to binary (base 2)
	result, err := convert(42, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%d converted to base %d is %s\n", 42, 2, result) // Output: 42 converted to base 2 is 101010
	}

	// Example: Convert decimal 1234 to hexadecimal (base 16)
	result, err = convert(1234, 16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%d converted to base %d is %s\n", 1234, 16, result) // Output: 1234 converted to base 16 is 4D2
	}

	// Example: Try to convert decimal 1234 to base 1 (not valid)
	result, err = convert(1234, 1)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: base must be between 2 and 36
	} else {
		fmt.Printf("%d converted to base %d is %s\n", 1234, 1, result)
	}
}
