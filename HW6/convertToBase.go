package main

import (
	"fmt"
	"strconv"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func convert(a, base int) (string, error) {
	result := ""
	dig := 0
	var err error
	if base > 73 {
		err.Error()
		return result, err
	}
	for a > 0 {
		dig = a % base
		if dig < 10 {
			result += strconv.Itoa(dig)
		} else {
			result += string(rune('A' + dig - 10))
		}
		if a < base {
			break
		}
		a /= base
	}
	return reverse(result), nil
}

func main() {
	// expected - 1111
	s1, _ := convert(15, 2)
	fmt.Println(s1)
	// expected - FF
	s2, _ := convert(255, 16)
	fmt.Println(s2)
	// expected - ~
	s3, _ := convert(71, 72)
	fmt.Println(s3)
}
