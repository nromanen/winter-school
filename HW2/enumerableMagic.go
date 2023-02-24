package main

import "fmt"

func eachCons(s []int, n int) (sub [][]int) {
	for i := 0; i < len(s)-n+1; i++ {
		sub = append(sub, s[i:i+n])
	}
	return
}

func main() {
	fmt.Println(eachCons([]int{1, 2, 3, 4}, 1))
	fmt.Println(eachCons([]int{1, 2, 3, 4}, 2))
	fmt.Println(eachCons([]int{1, 2, 3, 4}, 3))
}
