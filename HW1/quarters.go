package main

import "fmt"

func main() {

	var month int

	fmt.Println("Please, enter month:")
	fmt.Scanf("%d\n", &month)


	for month < 1 || month > 12 {
		fmt.Println("Please, enter correct number of month:")
		fmt.Scanf("%d\n", &month)
	}

	if month <= 3 {
		fmt.Println("it's 1")
	} else if month <= 6 {
		fmt.Println("it's 2")
	} else if month <= 9 {
		fmt.Println("it's 3")
	} else {
		fmt.Println("it's 4")
	}

}
