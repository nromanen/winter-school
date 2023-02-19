package main

import (
	"fmt"
)

func evenOrOdd() {
	var num int
	fmt.Print("Enter num : ")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println(num, " is Even")
	} else {
		fmt.Println(num, " is Odd")
	}

}

func indexW() {

	var m float32
	var h float32

	fmt.Println("enter your height in cm: ")
	fmt.Scan(&h)
	h = h / 100
	fmt.Println("enter your mass in kg: ")
	fmt.Scan(&m)

	var res = (m / (h * h))
	fmt.Print("Your body mass index is: ", res)

	switch {
	case res < 16.0:
		fmt.Println(" Underweight (Severe thinness)")
	case res < 17:
		fmt.Println(" (Underweight (Moderate thinness)")
	case res < 18.5:
		fmt.Println(" (Underweight (Mild thinness))")
	case res < 25:
		fmt.Println(" (Normal range)")
	case res < 30:
		fmt.Println(" (Overweight (Pre-obese))")
	case res < 35:
		fmt.Println(" (Obese (Class I))")
	case res < 40:
		fmt.Println(" (Obese (Class II))")
	case res >= 40:
		fmt.Println(" (Obese (Class III))")
	default:
		fmt.Println("You died -_-")
	}

}

func quarters() {
	var month int
	fmt.Print(" Enter number of month:")
	for month < 1 || month > 12 {
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

func main() {
	fmt.Println("-------------EX1------------- ")
	evenOrOdd()

	fmt.Println("-------------EX2-------------")
	indexW()

	fmt.Println("-------------EX3-------------")
	quarters()
}
