package main

import (
	"fmt"
)

func main() {
	var m int
	var h float64

	fmt.Println("Enter your body weight in kilograms:")
	fmt.Scan(&m)
	fmt.Println("Enter your height in meters:")
	fmt.Scan(&h)

	i := float64(m) / (h * h)
	fmt.Print("Your category: ")
	switch {
	case i < 18.5:
		fmt.Println("Insufficient mass")
	case i <= 24.9:
		fmt.Println("Norm")
	case i <= 29.9:
		fmt.Println("Obesity (smoothness)")
	case i <= 34.9:
		fmt.Println("Obesity of the first degree")
	case i <= 39.9:
		fmt.Println("Obesity of the second degree")
	case i >= 40:
		fmt.Println("Obesity of the third degree")
	
	default:
		fmt.Println("ERROR")
	}
}