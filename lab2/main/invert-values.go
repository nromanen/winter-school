package main

import "fmt"

func invertValues(s []int) []int {
	for i := range s {
		s[i] = -s[i]
	}
	return s
}

func main() {
	s := [][]int{
		{1, 2, 3, 4, 5},
		{1, -2, 3, -4, 5},
		{},
	}
	for _, v := range s {
		fmt.Println(invertValues(v))
	}
}
