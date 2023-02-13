package main

import "fmt"

func main() {
	var month int
	fmt.Print("Enter month number: ")
	fmt.Scan(&month)
	for {
		if month != 0 {
			break
		}
		fmt.Print("0 is not a month number, choose actual month number: ")
		fmt.Scan(&month)
	}

	if 1 < month && month < 4 {
		fmt.Println("It's first quater!")
	} else if 3 < month && month < 7 {
		fmt.Println("It's second quater!")
	} else if 6 < month && month < 10 {
		fmt.Println("It's third quater!")
	} else {
		fmt.Println("It's fourth quater!")
	}
}
