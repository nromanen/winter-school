package main

import "fmt"

func each_cons(list []int, n int) [][]int {
	var result [][]int
	for i := 0; i <= len(list)-n; i++ {
		result = append(result, list[i:i+n])
	}
	return result
}

func main() {
	var x int
	var y int
	fmt.Print("Input length: ")
	fmt.Scan(&x)
	var list = []int{}
	for i := 0; i < x; i++ {
		fmt.Printf("Input %v number: ", i+1)
		fmt.Scan(&y)
		list = append(list, y)
	}
	var z int
	fmt.Print("Input number: ")
	fmt.Scan(&z)
	fmt.Print(each_cons(list, z))
}