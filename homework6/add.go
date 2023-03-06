package main

import "fmt"

func Add(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func main() {

	f1 := Add(2)
	f2 := Add(10)

	result1 := f1(5) // result1 = 7
	result2 := f2(5) // result2 = 15

	fmt.Println(result1)
	fmt.Println(result2)

}
