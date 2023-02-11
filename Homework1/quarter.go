package main

import "fmt"

func main() {
	var month int
	fmt.Println("Input month as a number from 1 to 12:")
	fmt.Scan(&month)
	switch {
	case month > 0 && month < 4:
		fmt.Println("This month is in the first quarter")
	case month > 3 && month < 7:
		fmt.Println("This month is in the second quarter")
	case month > 6 && month < 10:
		fmt.Println("This month is in the third quarter")
	case month > 9 && month < 13:
		fmt.Println("This month is in the fourth quarter")
	default:
		fmt.Println("Error!")
	}
}
