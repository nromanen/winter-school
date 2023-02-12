package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d\n", &number)
	if number <= 3 && number >= 1 {
		fmt.Printf("first quarter")
	} else if number > 3 && number <= 6 {
		fmt.Printf("second quarter")
	} else if number > 6 && number <= 9 {
		fmt.Printf("third quarter")
	} else if number > 9 && number <= 12 {
		fmt.Printf("fourth quarter")
	}
}
