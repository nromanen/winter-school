package main

import "fmt"

func main() {

	var mount int16
	fmt.Println("Input month as a number from 1 to 12:")
	fmt.Scan(&mount)

	switch {
	case mount >= 1 && mount <= 3:
		fmt.Println("This month is in the first quarter")
	case mount >= 4 && mount <= 6:
		fmt.Println("This month is in the second quarter")
	case mount >= 7 && mount <= 9:
		fmt.Println("This month is in the thirt quarter")
	case mount >= 10 && mount <= 12:
		fmt.Println("This month is in the fourth quarter")

	default:
		fmt.Print("Invalid mount")

	}
}