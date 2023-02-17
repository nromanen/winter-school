package main

import "fmt"

func invertValues(l []int) []int {
	for i := range s {
		l[i] = -l[i]
	}
	return l
}

func main() {
	fmt.Println(invertValues([]int{1, 2, 3, 4, 5}))
	fmt.Println(invertValues([]int{1, -2, 3, -4, 5}))
	fmt.Println(invertValues([]int{}))
}