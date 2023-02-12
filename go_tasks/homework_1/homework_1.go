package main

import (
	"fmt"
)

func evenOrOdd() {
	var number int
	fmt.Print("Enter number > ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}
}

func yearQuarter() {
	var number int = 0
	for number <= 0 || number > 12 {
		fmt.Print("Enter number > ")
		fmt.Scan(&number)
		switch {
		case number <= 0:
			fmt.Sprintln("Incorrect month! (<0) Try again!")
		case number <= 3:
			fmt.Println("first quarter")
		case number <= 6:
			fmt.Println("second quarter")
		case number <= 9:
			fmt.Println("third quarter")
		case number <= 12:
			fmt.Println("fourth quarter")
		default:
			fmt.Println("Incorrect month! (>12) Try again!")
		}
	}
}

func indexWeight() {
	var weight float64
	var height float64
	fmt.Print("Enter weight in kilograms > ")
	fmt.Scan(&weight)
	fmt.Print("Enter height in meters > ")
	fmt.Scan(&height)
	var index = (weight / (height * height))
	switch {
	case index < 18.5:
		fmt.Println("insufficient mass")
	case index <= 24.9:
		fmt.Println("norm")
	case index < 29.9:
		fmt.Println("obesity")
	case index < 34.9:
		fmt.Println("obesity of 1st degree")
	case index < 39.9:
		fmt.Println("obesity of 2st degree")
	case index > 40:
		fmt.Println("obesity of 3st degree")
	default:
		fmt.Println("Error!")
	}
}

func main() {
	evenOrOdd()
	yearQuarter()
	indexWeight()
}
