package main

import "fmt"

func invertValues(s []int) []int {
	for i := range s {
		s[i] = -s[i]
	}
	return s
}

func main() {
	fmt.Println(invertValues([]int{1, 2, 3, 4, 5}))
	fmt.Println(invertValues([]int{1, -2, 3, -4, 5}))
	fmt.Println(invertValues([]int{}))
}
