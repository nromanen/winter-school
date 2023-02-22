package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CreatePhoneNumber(s []int) (string) {
	var format = "(XXX) XXX-XXXX"
	for i := 0; i < 10; i++ {
		n := s[i]
		format = strings.Replace(format, "X", strconv.Itoa(n), 1)
	//	fmt.Println(format)
	}
	return format
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(CreatePhoneNumber(x))

}

// CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
