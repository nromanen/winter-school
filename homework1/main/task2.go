package main

import "fmt"

func main() {
	var num int
	fmt.Println("Input number of month")
	fmt.Scan(&num)
	switch num {
	case 1, 2, 3:
		fmt.Println("Month #", num, "is in the first quarter")
	case 4, 5, 6:
		fmt.Println("Month #", num, "is in the second quarter")
	case 7, 8, 9:
		fmt.Println("Month #", num, "is in the third quarter")
	case 10, 11, 12:
		fmt.Println("Month #", num, "is in the fourth quarter")
	default:
		fmt.Println("Its not month number, u stupid")
	}
}
