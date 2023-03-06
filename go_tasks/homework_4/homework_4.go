package main

import (
	"fmt"
	"strings"
)

func IsDivisible(n, x, y int) {
	fmt.Println((n%x == 0) && (n%y == 0))
}

// used for Solve function
func checkIsElemIn(elem int, elements []int) bool {

	for i := 0; i < len(elements); i++ {

		if elem == elements[i] {
			return false
		}
	}
	return true
}

func Solve(arr []int) []int {
	elements := []int{arr[0]}

	for i := 0; i < len(arr); i++ {
		if checkIsElemIn(arr[i], elements) {
			elements = append(elements, arr[i])
		}
	}

	//fmt.Println(elements)

	var frequencies [][]int

	for i := 0; i < len(elements); i++ {
		frequency := 0
		for j := 0; j < len(arr); j++ {
			if arr[j] == elements[i] {
				frequency++
			}
		}
		frequencies = append(frequencies, []int{elements[i], frequency})
	}

	//fmt.Println(frequencies)

	for i := 0; i < len(frequencies)-1; i++ {
		for j := 0; j < len(frequencies)-1; j++ {
			if frequencies[j][1] < frequencies[j+1][1] {
				temp := frequencies[j]
				frequencies[j] = frequencies[j+1]
				frequencies[j+1] = temp
			}
		}
	}

	//fmt.Println(frequencies)

	var result []int

	for i := 0; i < len(frequencies); i++ {
		for j := 0; j < frequencies[i][1]; j++ {
			result = append(result, frequencies[i][0])
		}
	}

	//fmt.Println(result)

	for i := 0; i < len(result); i++ {
		arr[i] = result[i]
	}

	return arr
}

func InArray(array1 []string, array2 []string) {
	fmt.Println(array1)
	fmt.Println(array2)
	var result []string
	for i := 0; i < len(array2); i++ {
		for j := 0; j < len(array1); j++ {
			array1_word := array1[j]
			array2_word := array2[i]
			if strings.Contains(array2_word, array1_word) {

				result = append(result, array2_word)
			}
		}
	}
	fmt.Println(result)
}

func main() {
	//IsDivisible(3, 3, 4)
	//fmt.Println(Solve([]int{2, 3, 5, 3, 7, 9, 5, 3, 7}))
	//InArray([]string{"live", "arp", "strong"}, []string{"lively", "alive", "harp", "sharp", "armstrong"})
}
