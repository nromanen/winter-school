// https://www.codewars.com/kata/55fd2d567d94ac3bc9000064
package main

import "fmt"
import "math"

func RowSumOddNumbers(n int) int {
	res := math.Pow(float64(n), 3)
	fmt.Print(res)
	return int(res)
}
