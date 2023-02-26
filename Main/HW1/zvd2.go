package main

import "fmt"

func main() {
    var number int
    fmt.Print("Enter number: ")
    fmt.Scan(&number)
    quarter := (number-1)/3 + 1
    fmt.Printf("̳Mounth %d belongs to %d the quarter of the year..\n", number, quarter)
}