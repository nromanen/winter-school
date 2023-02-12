package main

import (
	"fmt"
	"math"
)

func main() {
	var m float64
	var h float64

	fmt.Println("Input your body mass:")
	fmt.Scan(&m)
	fmt.Println("Input your height in meters:")
	fmt.Scan(&h)

	index := (m) / (h * h)
	fmt.Print("Your BMI category: ")
	switch {
	case index > 0 && index < 18.5:
		fmt.Println("Insufficient mass")
	case index >= 18.5 && index < 24.9:
		fmt.Println("Normal range")
	case index >= 25.0 && index < 29.9:
		fmt.Println("Obesity (smoothness)")
	case index >= 30.0 && index < 34.9:
		fmt.Println("Obesity of the first degree")
	case index >= 35.0 && index < 39.9:
		fmt.Println("Obesity of the second degree")
	case index >= 40.0 && !math.IsInf(index, 0):
		fmt.Println("Obesity of the third degree")
	default:
		fmt.Println("Invalid input")
	}
}
