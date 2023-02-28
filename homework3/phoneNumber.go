package main

import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
	formatted := "(%d%d%d) %d%d%d-%d%d%d%d"
	return fmt.Sprintf(formatted, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

func main() {

	numbers := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(CreatePhoneNumber(numbers)) // prints "(123) 456-7890"

}
