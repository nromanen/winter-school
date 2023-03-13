package main

import "fmt"

func GetSubArrays(arr []int, subArrSize int) [][]int {
	var result [][]int
	for i := 0; i < len(arr)-subArrSize+1; i++ {
		result = append(result, arr[i:subArrSize+i])
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4}
	subArrSize := 2
	fmt.Printf("Subarrays of size %d in %v: %v", subArrSize, arr, GetSubArrays(arr, subArrSize))
}
