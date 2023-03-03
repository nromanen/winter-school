package main
// task 1
import (
	"errors"
	"strconv"
)

func convert(a, base int) (string, error) {
	if base < 2 || base > 36 {
		return "", errors.New("invalid base")
	}
	if a == 0 {
		return "0", nil
	}

	digits := "0123456789abcdefghijklmnopqrstuvwxyz"
	sign := ""
	if a < 0 {
		sign = "-"
		a = -a
	}

	var result string
	for a != 0 {
		remainder := a % base
		result = string(digits[remainder]) + result
		a /= base
	}

	return sign + result, nil
}
