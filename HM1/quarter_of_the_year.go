package main

import (
	"fmt"
)

func main() {
	var month uint
	fmt.Println("Input month as a number from 1 to 12:")
	fmt.Scan(&month)
	switch {
	case month >= 1 && month <= 3:
		fmt.Println("This month is in the first quarter")
	case month >= 4 && month <= 6:
		fmt.Println("This month is in the second quarter")
	case month >= 7 && month <= 9:
		fmt.Println("This month is in the third quarter")
	case month >= 10 && month <= 12:
		fmt.Println("This month is in the fourth quarter")
	default:
		fmt.Println("Invalid month")
	}
}
