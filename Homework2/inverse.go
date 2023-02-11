package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Print("Input length: ")
	fmt.Scan(&x)
	var list = []int{}
	for i := 0; i < x; i++ {
		fmt.Printf("Input %v number:", i+1)
		fmt.Scan(&y)
		list = append(list, y)
	}
	var result []int
	for i := 0; i < len(list); i++ {
		if list[i] == 0 {
			result = append(result, list[i])
		} else {
			result = append(result, list[i]*-1)
		}
	}
	fmt.Print(result)
}
