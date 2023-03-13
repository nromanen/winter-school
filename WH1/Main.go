package main

import "fmt"

func main() {
    // Завдання 1
    var number int
    fmt.Println("Enter a number:")
    fmt.Scan(&number)
    if number%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }

    // Завдання 2
    var month int
    fmt.Println("Enter a month (1-12):")
    fmt.Scan(&month)
    switch {
    case month >= 1 && month <= 3:
        fmt.Println("First quarter")
    case month >= 4 && month <= 6:
        fmt.Println("Second quarter")
    case month >= 7 && month <= 9:
        fmt.Println("Third quarter")
    case month >= 10 && month <= 12:
        fmt.Println("Fourth quarter")
    default:
        fmt.Println("Invalid month")
    }

    // Завдання 3
    var m int
    var h float64
    fmt.Println("Enter weight (kg):")
    fmt.Scan(&m)
    fmt.Println("Enter height (m):")
    fmt.Scan(&h)
    bmi := float64(m) / (h * h)
    fmt.Println("BMI:", bmi)
    switch {
    case bmi < 18.5:
        fmt.Println("Underweight")
    case bmi >= 18.5 && bmi < 25:
        fmt.Println("Normal weight")
    case bmi >= 25 && bmi < 30:
        fmt.Println("Overweight")
    case bmi >= 30:
        fmt.Println("Obesity")
    }
}
