package main

import "fmt"

func main() {
	var month int
	fmt.Println("Input month as a number 1-12: ")
	fmt.Scan(&month)
	switch month {
	case 1, 2, 3:
		fmt.Println("This month is part of the first quarter")
	case 4, 5, 6:
		fmt.Println("This month is part of the second quarter")
	case 7, 8, 9:
		fmt.Println("This month is part of the third quarter")
	case 10, 11, 12:
		fmt.Println("This month is part of the fourth  quarter")
	default:
		fmt.Println("Input valid month number")
	}
}
