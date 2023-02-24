package main

import (
	"fmt"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	result := fmt.Sprintf("%v", numbers)
	result = strings.ReplaceAll(result[1:len(result)-1], " ", "")
	return "(" + result[:3] + ") " + result[3:6] + "-" + result[6:]
}

func main() {

	fmt.Println(CreatePhoneNumber([10]uint{3, 8, 0, 6, 6, 5, 7, 5, 4, 3}))
}
