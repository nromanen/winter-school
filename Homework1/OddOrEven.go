package main

import "fmt"

func NumberIsOddOrEven(a int) bool {
    var status bool 
    status = false
    if (a%2 == 0) {
        status = true
    }
    return status
}

func main() {
    fmt.Print("Enter number : ")
    var n int
    fmt.Scanln(&n)

    if NumberIsOddOrEven(n) {
        fmt.Println(n, "is Even number")
    } else {
        fmt.Println(n, "is Odd number")
    }
}
