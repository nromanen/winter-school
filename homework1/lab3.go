package main

import (
	"fmt"
)

func main() {
	var mass uint
	var height float64

	fmt.Println("input you mass:")
	fmt.Scan(&mass)
	fmt.Println("input you heigth:")
	fmt.Scan(&height)

	index := float64(mass) / (height * height)
	fmt.Println("Your BMI category:")

	switch {
	case index < 18.5:
		fmt.Println("Insufficient mass")
	case index >= 18.5 && index < 25:
		fmt.Println("Norm")
	case index >= 25 && index < 30:
		fmt.Println("Obesity")
	case index >= 30 && index < 35:
		fmt.Println("Obesity of the first degree")
	case index >= 35 && index < 40:
		fmt.Println("Obesity of the second degree")
	case index >= 40:
		fmt.Println("Obesity of the third degree")

	default:
		fmt.Println("Error")

	}

}
