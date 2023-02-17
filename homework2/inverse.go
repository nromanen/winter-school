package main

import "fmt"

func invert(array []int) []int {
	for i := range array {
		array[i] = -array[i]
	}
	return array
}

func main() {

	fmt.Println(invert([]int{1, 2, 3, 4, 5}))
	fmt.Println(invert([]int{1, -2, 3, -4, 5}))
	fmt.Println(invert([]int{}))

}
