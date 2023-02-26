package main

import (
    "fmt"
)

func main() {
    fmt.Print("Please enter first number: ")
    var number int
    fmt.Scanln( & number) // take input from user
    if number % 2 == 0 {
        fmt.Printf("%v is a Even Number", number)
    } else {
        fmt.Printf("%v is a Odd Number", number)
    }
}