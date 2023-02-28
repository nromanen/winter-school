package main

import (
	"fmt"
	"math"
)

func main() {
	var m int
	var h float64

	fmt.Println("Input your body mass(in kilograms):")
	fmt.Scan(&m)
	fmt.Println("Input your height(in meters):")
	fmt.Scan(&h)

	index := float64(m) / (h * h)
	fmt.Print("Your BMI category: ")
	switch {
	case index >= 0 && index < 18.5:
		fmt.Println("Underweight")
	case index >= 18.5 && index < 25.0:
		fmt.Println("Normal")
	case index >= 25.0 && index < 30.0:
		fmt.Println("Overweight")
	case index >= 30.0 && index < 35.0:
		fmt.Println("Obese (Class I)")
	case index >= 35.0 && index < 40.0:
		fmt.Println("Obese (Class II)")
	case index >= 40.0 && !math.IsInf(index, 0):
		fmt.Println("Obese (Class III)")
	default:
		fmt.Println("smile!")
	}
}