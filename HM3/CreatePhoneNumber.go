package kata

import "fmt"

//https://www.codewars.com/kata/525f50e3b73515a6db000b83/train/go

func CreatePhoneNumber(numbers [10]uint) string {
	s := ""
	for _, n := range numbers {
		s += fmt.Sprintf("%v", n)
	}
	return fmt.Sprintf("(%v) %v-%v", s[0:3], s[3:6], s[6:10])
}
