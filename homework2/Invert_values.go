package main

import "fmt"

func invertValues(inVal []int) []int {

	for i := range inVal {
		inVal[i] = -inVal[i]
	}

	return inVal
}
func main() {
	
	fmt.Println(invertValues([]int{1, 2, 3, 4, 5}))

	fmt.Println(invertValues([]int{1, -2, 3, -4, 5}))

	fmt.Println(invertValues([]int{}))
}
