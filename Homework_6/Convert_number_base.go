package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := convert(123, 16)
	if err != nil {
		fmt.Println("Something happened wrong")
	}
	fmt.Println(result)
}

func convert(num, base int) (string, error) {
	if base < 2 || base > 36 {
		return "", errors.New("invalid base")
	}
	const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	var result []byte
	for num > 0 {
		rem := num % base
		result = append([]byte{charset[rem]}, result...)
		num /= base
		//fmt.Println(num)
	}
	if len(result) == 0 {
		result = []byte{'0'}
	}
	return string(result), nil
}
