package main

import (
	"fmt"
	"math"
)

func main() {
	var m uint
	var h float64

	fmt.Println("Input your body mass:")
	fmt.Scan(&m)
	fmt.Println("Input your height in meters:")
	fmt.Scan(&h)

	index := float64(m) / (h * h)
	fmt.Print("Your body mass index category: ")
	switch {
	case index > 0 && index < 16.0:
		fmt.Println("Underweight (Severe thinness)")
	case index >= 16.0 && index < 17.0:
		fmt.Println("Underweight (Moderate thinness)")
	case index >= 17.0 && index < 18.5:
		fmt.Println("Underweight (Mild thinness)")
	case index >= 18.5 && index < 25.0:
		fmt.Println("Normal range")
	case index >= 25.0 && index < 30.0:
		fmt.Println("Overweight (Pre-obese)")
	case index >= 30.0 && index < 35.0:
		fmt.Println("Obese (Class I)")
	case index >= 35.0 && index < 40.0:
		fmt.Println("Obese (Class II)")
	case index >= 40.0 && !math.IsInf(index, 0):
		fmt.Println("Obese (Class III)")
	default:
		fmt.Println("Invalid input")
	}
}
