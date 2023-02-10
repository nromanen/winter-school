package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Println("Please enter month as a number from 1 to 12:")
	fmt.Scan(&m)
	switch {
	case m <= 3:
		fmt.Println("This month is in the first quarter")
	case m <= 6:
		fmt.Println("This month is in the second quarter")
	case m <= 9:
		fmt.Println("This month is in the third quarter")
	case m <= 12:
		fmt.Println("This month is in the fourth quarter")
	default:
		fmt.Println("Error")
	}
}

