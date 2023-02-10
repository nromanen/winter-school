package main

import (
	"fmt"
)

//Task 5 https://www.codewars.com/kata/59a8570b570190d313000037/train/go
func sumCubes(n int) int {

   var i, sum int

   sum = 0

   for i = 1; i <= n; i = i + 1 {

      sum = sum + i * i * i
   }

   return sum
}
func main() {

    var n, answer int

    fmt.Println("Sums all the cubed values from 1 to n (inclusive), and returns that sum.")
    fmt.Println("Please enter positive integer n: ")
    fmt.Scan(&n)
	
    answer = sumCubes(n)
    fmt.Println("Sum of the cubes", n, "is", answer)
   
}
   