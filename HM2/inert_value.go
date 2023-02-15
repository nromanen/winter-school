package main
//https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad/go
import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i] * -1
	}
	fmt.Print(numbers)
}
