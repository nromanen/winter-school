package main

import "fmt"

func createPhoneNumber(numbers []int) string {
	return fmt.Sprintf("(%v%v%v) %v%v%v-%v%v%v%v", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

func main() {
	var y int
	var numbers = []int{}
	fmt.Printf("Input 10 numbers: ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&y)
		numbers = append(numbers, y)
	}
	fmt.Print(createPhoneNumber(numbers))

}
