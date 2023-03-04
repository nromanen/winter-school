package main

import (
	"fmt"
)

func CreatePhoneNumber(numbers []uint) string {
	var str string
	for _, n := range numbers {
		str += string('0' + n)
	}
	return fmt.Sprintf("(%v) %v-%v", str[:3], str[3:6], str[6:])
}
func main() {

	fmt.Println("Task 5 :", CreatePhoneNumber([]uint{3, 7, 6, 7, 1, 2, 3, 4, 9, 0}))

}
