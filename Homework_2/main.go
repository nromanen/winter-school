package main

import (
	"fmt"
)

func Invert(arr []int) []int {
	var res []int = make([]int, len(arr))

	for i, v := range arr {
		res[i] = -v
	}

	return res
}
func RowSumOddNumbers(n int) int {
	return n * n * n
}

func monkeyCount(n int) []int {
	var a = make([]int, n)

	for i := 1; i < n+1; i++ {
		a[i-1] = i
	}

	return a
}

func EachCons(arr []int, n int) (res [][]int) {
	for i := 0; i <= len(arr)-n; i++ {
		var tmp []int
		for j := 0; j < n; j++ {
			tmp = append(tmp, arr[i+j])
		}
		res = append(res, tmp)
	}
	return
}

func SumCubes(n int) int {
	if n == 1 {
		return 1
	} else {
		return n*n*n + SumCubes(n-1)
	}
}
func main() {

	fmt.Println("Task 1 :", Invert([]int{1, -2, 3, 4, 5}))
	fmt.Println("Task 2 :", RowSumOddNumbers(10))
	fmt.Println("Task 3 :", monkeyCount(6))
	fmt.Println("Task 4 :", EachCons([]int{1, 2, 3, 4, 5, 6}, 2))
	fmt.Println("Task 5 :", SumCubes(3))
}
