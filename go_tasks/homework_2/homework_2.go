package main

import "fmt"

func Invert(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = (arr[i] * -1)
	}
	return arr
}

func RowSumOddNumbers(n int) int {
	return n * n * n
}

func monkeyCount(n int) []int {

	result := []int{}
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}
	return result
}

func EachCons(arr []int, n int) [][]int {
	var result = [][]int{}
	for i := 0; i <= (len(arr) - n); i++ {
		slice := arr[i : i+n]
		result = append(result, slice)
	}
	return result
}

func SumCubes(n int) int {
	var result int = 0
	for i := 1; i <= n; i++ {
		result += (i * i * i)
	}
	return result
}

func main() {
	fmt.Println("Task 1 Invert", Invert([]int{1, -3, 4, 5}))
	fmt.Println("Task 2 RowSumOddNumbers", RowSumOddNumbers(42))
	fmt.Println("Task 3 monkeyCount", monkeyCount(42))
	fmt.Println("Task 4 EachCons", EachCons([]int{3, 5, 8, 13}, 2))
	fmt.Println("Task 5 SumCubes", SumCubes(3))
}
